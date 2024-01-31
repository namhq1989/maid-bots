package auth

import (
	"github.com/namhq1989/maid-bots/pkg/sso"
	"github.com/namhq1989/maid-bots/util/appcontext"
	jwt2 "github.com/namhq1989/maid-bots/util/jwt"
)

// SSO ...
type SSO struct{}

func (SSO) LoginWithGoogle(ctx *appcontext.AppContext, token string) (string, error) {
	ctx.Logger.Text("new Google sso")
	var (
		userSvc = User{}
	)

	// fetch data from google
	googleData, err := sso.LoginWithGoogle(ctx, token)
	if err != nil {
		return "", err
	}

	ctx.Logger.Text("fetched Google data successfully")

	// check user is existed or not
	user, _ := userSvc.FindByGoogleID(ctx, googleData.ID)
	if user == nil {
		ctx.Logger.Info("user with Google id not found, create one", appcontext.Fields{"googleData": googleData})
		user, err = userSvc.CreateWithGoogleData(ctx, *googleData)
		if err != nil {
			return "", err
		}
	} else {
		ctx.Logger.Text("user found")
	}

	ctx.Logger.Text("generate jwt token")

	// generate token and return
	return jwt2.Signing(ctx, jwt2.User{ID: user.ID.String()})
}

func (SSO) LoginWithGitHub(ctx *appcontext.AppContext, code string) (string, error) {
	ctx.Logger.Text("new GitHub sso")
	var (
		userSvc = User{}
	)

	// fetch data from GitHub
	githubData, err := sso.LoginWithGitHub(ctx, code)
	if err != nil {
		return "", err
	}

	ctx.Logger.Text("fetched GitHub data successfully")

	// check user is existed or not
	user, _ := userSvc.FindByGitHubID(ctx, githubData.ID)
	if user == nil {
		ctx.Logger.Info("user with GitHub id not found, create one", appcontext.Fields{"githubData": githubData})
		user, err = userSvc.CreateWithGitHubData(ctx, *githubData)
		if err != nil {
			return "", err
		}
	}

	ctx.Logger.Text("generate jwt token")

	// generate token and return
	return jwt2.Signing(ctx, jwt2.User{ID: user.ID.String()})
}
