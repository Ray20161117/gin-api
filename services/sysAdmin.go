/**
 * 用户 服务层
 */
package services

import (
	"gin-api/common/response"
	"gin-api/common/utils"
	"gin-api/models/dto"
	"gin-api/models/entity"
	"gin-api/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// 定义接口
type ISysAdminService interface {
	// 登录
	Login(c *gin.Context, loginDto entity.LoginDto)
	// 新增
	//CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	// 根据ID查询详情
	//GetSysAdminInfo(c *gin.Context, Id int)
	// 编辑
	//UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto)
	// 删除
	//DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto)
	// 修改状态
	//UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto)
	// 重置密码
	//ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto)
	// 获取列表
	//GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string)
	// 修改跟人信息--个人中心
	//UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalDto)
	// 修改个人密码
	//UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto)
}

type SysAdminServiceImpl struct {
	validator *validator.Validate
}

// 初始化SysAdminServiceImpl
func NewSysAdminServiceImpl() *SysAdminServiceImpl {
	return &SysAdminServiceImpl{
		validator: validator.New(),
	}
}

// 登录
func (s SysAdminServiceImpl) Login(c *gin.Context, loginDto entity.LoginDto) {
	// 登录参数校验
	if err := s.validator.Struct(loginDto); err != nil {
		response.Failed(c, int(response.ApiCode.MissingLoginParameter), response.ApiCode.GetMessage(response.ApiCode.MissingLoginParameter))
		return
	}

	ip := c.ClientIP()

	// 验证码是否过期
	code := utils.RedisStore{}.Get(loginDto.IdKey, true)

	if len(code) == 0 {
		dto.CreateSysLoginInfo(loginDto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "验证码已过期", 2)
		response.Failed(c, int(response.ApiCode.VerificationCodeHasExpired), response.ApiCode.GetMessage(response.ApiCode.VerificationCodeHasExpired))
		return
	}

	// 校验验证码
	if !CaptVerify(loginDto.IdKey, loginDto.Captcha) {
		dto.CreateSysLoginInfo(loginDto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "验证码不正确", 2)
		response.Failed(c, int(response.ApiCode.CAPTCHANOTTRUE), response.ApiCode.GetMessage(response.ApiCode.CAPTCHANOTTRUE))
		return
	}

	// 校验用户信息
	sysAdmin, err := dto.SysAdminDetail(loginDto)
	if err != nil {
		response.Failed(c, int(response.ApiCode.QUERYUSERFAILED), response.ApiCode.GetMessage(response.ApiCode.QUERYUSERFAILED))
		return
	}

	if sysAdmin.Password != utils.EncryptionMd5(loginDto.Password) {
		dto.CreateSysLoginInfo(loginDto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "密码不正确", 2)
		response.Failed(c, int(response.ApiCode.PASSWORDNOTTRUE), response.ApiCode.GetMessage(response.ApiCode.PASSWORDNOTTRUE))
		return
	}

	if sysAdmin.Status == 2 {
		dto.CreateSysLoginInfo(loginDto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "账号已停用", 2)
		response.Failed(c, int(response.ApiCode.STATUSISENABLE), response.ApiCode.GetMessage(response.ApiCode.STATUSISENABLE))
		return
	}

	// 生成token
	tokenString, err := jwt.GenerateTokenByAdmin(sysAdmin)
	if err != nil {
		response.Failed(c, int(response.ApiCode.TOKENGENERATEFAILED), response.ApiCode.GetMessage(response.ApiCode.TOKENGENERATEFAILED))
		return
	}

	dto.CreateSysLoginInfo(loginDto.Username, ip, utils.GetRealAddressByIP(ip), utils.GetBrowser(c), utils.GetOs(c), "登录成功", 1)

	// 查询左侧菜单列表
	var leftMenuVo []entity.LeftMenuVoDto
	leftMenuList, err := dto.QueryLeftMenuList(sysAdmin.ID)
	if err != nil {
		response.Failed(c, int(response.ApiCode.QUERYLEFTMENUFAILED), response.ApiCode.GetMessage(response.ApiCode.QUERYLEFTMENUFAILED))
		return
	}

	for _, value := range leftMenuList {
		menuSvoList, err := dto.QueryMenuVoList(sysAdmin.ID, value.Id)
		if err != nil {
			response.Failed(c, int(response.ApiCode.QUERYLEFTMENUFAILED), response.ApiCode.GetMessage(response.ApiCode.QUERYLEFTMENUFAILED))
			return
		}

		item := entity.LeftMenuVoDto{
			MenuSvoList: menuSvoList,
			Id:          value.Id,
			MenuName:    value.MenuName,
			Icon:        value.Icon,
			Url:         value.Url,
		}
		leftMenuVo = append(leftMenuVo, item)
	}

	// 查询权限列表
	permissionList, err := dto.QueryPermissionValueList(sysAdmin.ID)
	if err != nil {
		response.Failed(c, int(response.ApiCode.QUERYPERMISSIONFAILED), response.ApiCode.GetMessage(response.ApiCode.QUERYPERMISSIONFAILED))
		return
	}

	var stringList = make([]string, 0)
	for _, value := range permissionList {
		stringList = append(stringList, value.Value)
	}

	response.Success(c, map[string]interface{}{
		"token":          tokenString,
		"sysAdmin":       sysAdmin,
		"leftMenuList":   leftMenuVo,
		"permissionList": stringList,
	})
}

var sysAdminService = NewSysAdminServiceImpl()

func SysAdminService() ISysAdminService {
	return sysAdminService
}