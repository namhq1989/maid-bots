package monitor

import (
	"fmt"
	"strings"

	modelcommand "github.com/namhq1989/maid-bots/internal/model/command"

	"github.com/namhq1989/maid-bots/pkg/mongodb"

	modelresponse "github.com/namhq1989/maid-bots/internal/model/response"
	"github.com/namhq1989/maid-bots/internal/service"
	"github.com/namhq1989/maid-bots/util/appcommand"
	"github.com/namhq1989/maid-bots/util/appcontext"
)

type Register struct {
	Target   string
	Value    string
	Platform string
	ChatID   string
	User     modelcommand.User
}

func (c Register) Process(ctx *appcontext.AppContext) (string, error) {
	// check first
	check := Check{
		Target: c.Target,
		Value:  c.Value,
	}

	result, err := check.Process(ctx)
	if err != nil {
		return "", err
	}

	switch c.Target {
	case appcommand.MonitorTargets.Domain.Name:
		return c.domain(ctx, result)
	case appcommand.MonitorTargets.HTTP.Name:
		return c.http(ctx, result)
	case appcommand.MonitorTargets.TCP.Name:
		return c.tcp(ctx, result)
	case appcommand.MonitorTargets.ICMP.Name:
		return c.icmp(ctx, result)
	}

	return "", fmt.Errorf("target %s is not supported", c.Target)
}

func (c Register) domain(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	var (
		userSvc    = service.User{}
		monitorSvc = service.Monitor{}
	)

	// find user first
	user, err := userSvc.FindOrCreateWithPlatformID(ctx, c.Platform, c.ChatID, c.User)
	if err != nil {
		return "", err
	}

	// check target is existed or not
	if monitorSvc.IsTargetExisted(ctx, mongodb.MonitorTypeDomain, checkData.Name, user.ID) {
		return "", fmt.Errorf("target %s is already registered", checkData.Name)
	}

	// create monitor
	doc, err := monitorSvc.CreateDomain(ctx, checkData.Name, checkData.Scheme, user.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("target `%s` has been successfully registered with id `%s`", checkData.Name, strings.ToUpper(doc.Code)), nil
}

func (c Register) http(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	var (
		userSvc    = service.User{}
		monitorSvc = service.Monitor{}
	)

	// find user first
	user, err := userSvc.FindOrCreateWithPlatformID(ctx, c.Platform, c.ChatID, c.User)
	if err != nil {
		return "", err
	}

	// check target is existed or not
	if monitorSvc.IsTargetExisted(ctx, mongodb.MonitorTypeHTTP, checkData.Name, user.ID) {
		return "", fmt.Errorf("target %s is already registered", checkData.Name)
	}

	// create monitor
	doc, err := monitorSvc.CreateHTTP(ctx, checkData.Name, user.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("target `%s` has been successfully registered with id `%s`", checkData.Name, strings.ToUpper(doc.Code)), nil
}

func (c Register) tcp(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	var (
		userSvc    = service.User{}
		monitorSvc = service.Monitor{}
	)

	// find user first
	user, err := userSvc.FindOrCreateWithPlatformID(ctx, c.Platform, c.ChatID, c.User)
	if err != nil {
		return "", err
	}

	// check target is existed or not
	if monitorSvc.IsTargetExisted(ctx, mongodb.MonitorTypeTCP, checkData.Name, user.ID) {
		return "", fmt.Errorf("target %s is already registered", checkData.Name)
	}

	// create monitor
	doc, err := monitorSvc.CreateTCP(ctx, checkData.Name, user.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("target `%s` has been successfully registered with id `%s`", checkData.Name, strings.ToUpper(doc.Code)), nil
}

func (c Register) icmp(ctx *appcontext.AppContext, checkData *modelresponse.Check) (string, error) {
	var (
		userSvc    = service.User{}
		monitorSvc = service.Monitor{}
	)

	// find user first
	user, err := userSvc.FindOrCreateWithPlatformID(ctx, c.Platform, c.ChatID, c.User)
	if err != nil {
		return "", err
	}

	// check target is existed or not
	if monitorSvc.IsTargetExisted(ctx, mongodb.MonitorTypeICMP, checkData.Name, user.ID) {
		return "", fmt.Errorf("target %s is already registered", checkData.Name)
	}

	// create monitor
	doc, err := monitorSvc.CreateICMP(ctx, checkData.Name, user.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("target `%s` has been successfully registered with id `%s`", checkData.Name, strings.ToUpper(doc.Code)), nil
}
