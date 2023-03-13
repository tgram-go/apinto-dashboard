package controller

import (
	"fmt"
	"github.com/eolinker/apinto-dashboard/access"
	"github.com/eolinker/apinto-dashboard/common"
	"github.com/eolinker/apinto-dashboard/dto"
	"github.com/eolinker/apinto-dashboard/enum"
	"github.com/eolinker/apinto-dashboard/service"
	"github.com/eolinker/eosc/common/bean"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type userController struct {
	userInfo service.IUserInfoService
}

func RegisterUserRouter(router gin.IRoutes) {
	u := &userController{}
	bean.Autowired(&u.userInfo)

	router.GET("/access", u.getAllAccess)
	router.GET("/my/modules", u.getUserAccess)
	router.PUT("/my/profile", logHandler(enum.LogOperateTypeEdit, enum.LogKindUser), u.updateMyProfile)
	router.GET("/my/profile", u.getMyProfile)
	router.POST("/my/password", logHandler(enum.LogOperateTypeEdit, enum.LogKindUser), u.updateMyPassword)

	router.GET("/roles", genAccessHandler(access.UserRoleView, access.UserRoleEdit), u.getRoleList)
	router.GET("/role", genAccessHandler(access.UserRoleView, access.UserRoleEdit), u.getRoleInfo)
	router.GET("/role/options", u.getRoleOptions)
	router.POST("/role", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeCreate, enum.LogKindRole), u.createRole)
	router.PUT("/role", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeEdit, enum.LogKindRole), u.updateRole)
	router.DELETE("/role", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeDelete, enum.LogKindRole), u.deleteRole)
	router.POST("/role/batch-update", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeEdit, enum.LogKindRole), u.roleBatchUpdate)
	router.POST("/role/batch-delete", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeEdit, enum.LogKindRole), u.roleBatchRemove)

	router.POST("/user/delete", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeDelete, enum.LogKindUser), u.delUser)
	router.POST("/user/profile", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeCreate, enum.LogKindUser), u.createUser)
	router.PATCH("/user/profile", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeEdit, enum.LogKindUser), u.patchUser)
	router.PUT("/user/profile", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeEdit, enum.LogKindUser), u.updateUser)
	router.GET("/user/profile", genAccessHandler(access.UserRoleView, access.UserRoleEdit), u.getUser)
	router.GET("/user/list", genAccessHandler(access.UserRoleView, access.UserRoleEdit), u.getUserList)
	router.GET("/user/enum", u.getUserList)
	router.POST("/user/password-reset", genAccessHandler(access.UserRoleEdit), logHandler(enum.LogOperateTypeEdit, enum.LogKindUser), u.resetUserPwd)
}

func (u *userController) getMyProfile(ginCtx *gin.Context) {
	userId := getUserId(ginCtx)

	userInfo, err := u.userInfo.GetUserInfo(ginCtx, userId)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("getMyProfile fail. err:%s", err.Error())))
		return
	}

	lastLogin := ""
	if userInfo.LastLoginTime != nil {
		lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
	}
	resUserInfo := dto.UserInfo{
		Id:           userInfo.Id,
		Sex:          userInfo.Sex,
		Avatar:       userInfo.Avatar,
		Desc:         userInfo.Remark,
		Email:        userInfo.Email,
		Phone:        userInfo.Phone,
		Status:       userInfo.Status,
		UserName:     userInfo.UserName,
		NickName:     userInfo.NickName,
		NoticeUserId: userInfo.NoticeUserId,
		LastLogin:    lastLogin,
		CreateTime:   common.TimeToStr(userInfo.CreateTime),
		UpdateTime:   common.TimeToStr(userInfo.UpdateTime),
		RoleIds:      strings.Split(userInfo.RoleIds, ","),
	}

	m := make(map[string]interface{})
	m["profile"] = resUserInfo
	m["describe"] = userInfo.Remark

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(m))
}

func (u *userController) updateMyProfile(ginCtx *gin.Context) {
	userId := getUserId(ginCtx)

	req := &dto.UpdateMyProfileReq{}
	err := ginCtx.BindJSON(req)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("updateMyProfile fail. err:%s", err.Error())))
		return
	}

	if err = u.userInfo.UpdateMyProfile(ginCtx, userId, req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("updateMyProfile fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))
}

func (u *userController) updateMyPassword(ginCtx *gin.Context) {
	userId := getUserId(ginCtx)

	req := &dto.UpdateMyPasswordReq{}
	err := ginCtx.BindJSON(req)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("updateMyPassword fail. err:%s", err.Error())))
		return
	}

	if err = u.userInfo.UpdateMyPassword(ginCtx, userId, req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("updateMyPassword fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))
}

func (u *userController) getAllAccess(ginCtx *gin.Context) {

	globalAccess, depth := access.GetGlobalAccessConfig()

	data := make(map[string]interface{})
	data["modules"] = getModules(globalAccess)
	data["depth"] = depth

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))

}

func getModules(modules []*access.GlobalAccess) []*dto.SystemModuleItem {
	items := make([]*dto.SystemModuleItem, len(modules))
	for i, module := range modules {
		item := &dto.SystemModuleItem{
			ID:     module.ID,
			Title:  module.Title,
			Module: module.Module,
			Access: module.Access,
		}
		if len(module.Children) > 0 {
			item.Children = getModules(module.Children)
		}
		items[i] = item
	}
	return items
}

func (u *userController) getUserAccess(ginCtx *gin.Context) {
	userID := getUserId(ginCtx)

	accessSet, err := u.userInfo.GetAccessInfo(ginCtx, userID)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("GetUserAccessList fail. err:%s", err.Error())))
		return
	}
	allModules := access.GetAllModulesConfig()

	modules := make([]*dto.UserModuleItem, 0, len(accessSet))
	for _, module := range allModules {
		for _, needId := range module.ModuleNeed {
			if _, has := accessSet[needId]; has {
				accessList := make([]string, 0, len(module.Access))
				for _, key := range module.Access {
					accessId, err := access.Parse(key)
					if err != nil {
						continue
					}
					if _, exist := accessSet[accessId]; exist {
						accessList = append(accessList, key)
					}
				}
				modules = append(modules, &dto.UserModuleItem{
					Id:     module.ID,
					Router: module.Router,
					Title:  module.Title,
					Access: accessList,
					Parent: module.Parent,
				})
				break
			}
		}
	}

	data := make(map[string]interface{})
	data["modules"] = modules

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))

}

func (u *userController) getRoleList(ginCtx *gin.Context) {
	userID := getUserId(ginCtx)
	roleList, totalUsers, err := u.userInfo.GetRoleList(ginCtx, userID)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("GetRoleList fail. err:%s", err.Error())))
		return
	}
	roles := make([]*dto.RoleListItem, 0, len(roleList))
	for _, item := range roleList {
		role := &dto.RoleListItem{
			ID:             item.ID,
			Title:          item.Title,
			UserNum:        item.UserNum,
			OperateDisable: item.OperateDisable,
			Type:           item.Type,
		}
		roles = append(roles, role)
	}

	data := make(map[string]interface{})
	data["roles"] = roles
	data["total"] = totalUsers
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))
}

func (u *userController) getRoleInfo(ginCtx *gin.Context) {
	roleID := ginCtx.Query("id")
	if roleID == "" {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("GetRoleInfo fail. err: id can't be nil")))
		return
	}

	info, err := u.userInfo.GetRoleInfo(ginCtx, roleID)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("GetRoleInfo fail. err:%s", err.Error())))
		return
	}
	role := &dto.ProxyRoleInfo{
		Title:  info.Title,
		Desc:   info.Desc,
		Access: info.Access,
	}
	data := make(map[string]interface{})
	data["role"] = role
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))
}

func (u *userController) getRoleOptions(ginCtx *gin.Context) {
	optionList, err := u.userInfo.GetRoleOptions(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("GetRoleOptions fail. err:%s", err.Error())))
		return
	}
	options := make([]*dto.RoleOptionItem, 0, len(optionList))
	for _, item := range optionList {
		option := &dto.RoleOptionItem{
			ID:             item.ID,
			Title:          item.Title,
			OperateDisable: item.OperateDisable,
		}
		options = append(options, option)
	}

	data := make(map[string]interface{})
	data["roles"] = options
	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(data))
}

func (u *userController) createRole(ginCtx *gin.Context) {
	userID := getUserId(ginCtx)

	input := new(dto.ProxyRoleInfo)
	if err := ginCtx.BindJSON(input); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}

	//Check Input
	if input.Title == "" {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("CreateRole fail. err: title can't be nil. ")))
		return
	}

	err := u.userInfo.CreateRole(ginCtx, userID, input)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("CreateAPI fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))
}

func (u *userController) updateRole(ginCtx *gin.Context) {
	userID := getUserId(ginCtx)

	roleID := ginCtx.Query("id")
	if roleID == "" {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("UpdateRole fail. err: id can't be nil")))
		return
	}

	input := new(dto.ProxyRoleInfo)
	if err := ginCtx.BindJSON(input); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}

	//Check Input
	if input.Title == "" {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("UpdateRole fail. err: title can't be nil. ")))
		return
	}

	err := u.userInfo.UpdateRole(ginCtx, userID, roleID, input)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("UpdateRole fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))
}

func (u *userController) deleteRole(ginCtx *gin.Context) {
	userID := getUserId(ginCtx)
	roleID := ginCtx.Query("id")
	if roleID == "" {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("DeleteRole fail. err: id can't be nil")))
		return
	}

	err := u.userInfo.DeleteRole(ginCtx, userID, roleID)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("DeleteRole fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))
}

func (u *userController) roleBatchUpdate(ginCtx *gin.Context) {
	input := new(dto.BatchUpdateRole)
	if err := ginCtx.BindJSON(input); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}

	err := u.userInfo.RoleBatchUpdate(ginCtx, input.Ids, input.RoleId)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("roleBatchUpdate fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))
}

func (u *userController) roleBatchRemove(ginCtx *gin.Context) {
	input := new(dto.BatchRemoveRole)
	if err := ginCtx.BindJSON(input); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}

	err := u.userInfo.RoleBatchRemove(ginCtx, input.Ids, input.RoleId)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("roleBatchRemove fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))
}

func (u *userController) delUser(ginCtx *gin.Context) {
	userID := getUserId(ginCtx)

	req := &dto.DelUserReq{}
	if err := ginCtx.BindJSON(req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("delUser fail. err:%s", err.Error())))
		return
	}

	if len(req.UserIds) == 0 {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("delUser fail. err:%s", "参数错误")))
		return
	}

	err := u.userInfo.DelUser(ginCtx, userID, req.UserIds)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("delUser fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}

func (u *userController) createUser(ginCtx *gin.Context) {
	userID := getUserId(ginCtx)

	req := &dto.SaveUserReq{}
	if err := ginCtx.BindJSON(req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("createUser fail. err:%s", err.Error())))
		return
	}

	if err := common.IsMatchString(common.EnglishOrNumber_, req.UserName); err != nil {
		return
	}

	if err := u.userInfo.CreateUser(ginCtx, userID, req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("createUser fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}

func (u *userController) patchUser(ginCtx *gin.Context) {
	operator := getUserId(ginCtx)
	userIDStr := ginCtx.Query("id")

	userId, err := strconv.Atoi(userIDStr)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("patchUser fail. err:%s", err.Error())))
		return
	}

	req := &dto.PatchUserReq{}
	if err = ginCtx.BindJSON(req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("patchUser fail. err:%s", err.Error())))
		return
	}

	if err = u.userInfo.PatchUser(ginCtx, operator, userId, req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("patchUser fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}

func (u *userController) updateUser(ginCtx *gin.Context) {
	operator := getUserId(ginCtx)

	userIdStr := ginCtx.Query("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("updateUser fail. err:%s", err.Error())))
		return
	}

	req := &dto.SaveUserReq{}
	if err := ginCtx.BindJSON(req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("updateUser fail. err:%s", err.Error())))
		return
	}

	if err = u.userInfo.UpdateUser(ginCtx, operator, userId, req); err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("updateUser fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}

func (u *userController) getUserList(ginCtx *gin.Context) {

	roleId := ginCtx.Query("role")
	keyword := ginCtx.Query("keyword")

	userInfoList, err := u.userInfo.GetUserInfoList(ginCtx, roleId, keyword)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("getUserList fail. err:%s", err.Error())))
		return
	}

	resList := make([]dto.UserInfo, 0, len(userInfoList))

	for _, userInfo := range userInfoList {
		lastLogin := ""
		if userInfo.LastLoginTime != nil {
			lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
		}
		resUserInfo := dto.UserInfo{
			Id:             userInfo.Id,
			Sex:            userInfo.Sex,
			Avatar:         userInfo.Avatar,
			Email:          userInfo.Email,
			Phone:          userInfo.Phone,
			Status:         userInfo.Status,
			UserName:       userInfo.UserName,
			NickName:       userInfo.NickName,
			NoticeUserId:   userInfo.NoticeUserId,
			LastLogin:      lastLogin,
			CreateTime:     common.TimeToStr(userInfo.CreateTime),
			UpdateTime:     common.TimeToStr(userInfo.UpdateTime),
			OperateDisable: userInfo.OperateEnable,
			Desc:           userInfo.Remark,
			Operator:       userInfo.Operator,
			RoleIds:        strings.Split(userInfo.RoleIds, ","),
		}
		resList = append(resList, resUserInfo)
	}

	m := make(map[string]interface{})
	m["users"] = resList

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(m))

}

func (u *userController) userEnum(ginCtx *gin.Context) {
	infos, err := u.userInfo.GetUserInfoAll(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(err.Error()))
		return
	}
	resList := make([]dto.UserInfo, 0, len(infos))
	for _, userInfo := range infos {
		resUserInfo := dto.UserInfo{
			Id:       userInfo.Id,
			Email:    userInfo.Email,
			UserName: userInfo.UserName,
			NickName: userInfo.NickName,
		}
		resList = append(resList, resUserInfo)
	}

	m := make(map[string]interface{})
	m["users"] = resList

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(m))
}

func (u *userController) getUser(ginCtx *gin.Context) {

	userIdStr := ginCtx.Query("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("getUser fail. err:%s", err.Error())))
		return
	}

	userInfo, err := u.userInfo.GetUserInfo(ginCtx, userId)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("getUser fail. err:%s", err.Error())))
		return
	}

	lastLogin := ""
	if userInfo.LastLoginTime != nil {
		lastLogin = common.TimeToStr(*userInfo.LastLoginTime)
	}
	resUserInfo := dto.UserInfo{
		Id:           userInfo.Id,
		Sex:          userInfo.Sex,
		Avatar:       userInfo.Avatar,
		Desc:         userInfo.Remark,
		Email:        userInfo.Email,
		NoticeUserId: userInfo.NoticeUserId,
		Phone:        userInfo.Phone,
		Status:       userInfo.Status,
		UserName:     userInfo.UserName,
		NickName:     userInfo.NickName,
		LastLogin:    lastLogin,
		CreateTime:   common.TimeToStr(userInfo.CreateTime),
		UpdateTime:   common.TimeToStr(userInfo.UpdateTime),
		RoleIds:      strings.Split(userInfo.RoleIds, ","),
	}

	m := make(map[string]interface{})
	m["profile"] = resUserInfo
	m["describe"] = userInfo.Remark

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(m))

}

func (u *userController) resetUserPwd(ginCtx *gin.Context) {

	operator := getUserId(ginCtx)

	resetPasswordReq := new(dto.ResetPasswordReq)
	err := ginCtx.BindJSON(resetPasswordReq)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("resetUserPwd fail. err:%s", err.Error())))
		return
	}

	err = u.userInfo.ResetUserPwd(ginCtx, operator, resetPasswordReq.Id, resetPasswordReq.Password)
	if err != nil {
		ginCtx.JSON(http.StatusOK, dto.NewErrorResult(fmt.Sprintf("resetUserPwd fail. err:%s", err.Error())))
		return
	}

	ginCtx.JSON(http.StatusOK, dto.NewSuccessResult(nil))

}
