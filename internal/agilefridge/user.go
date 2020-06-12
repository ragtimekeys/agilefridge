package agilefridge

import "time"

type (
	IdentifyingMetadata struct {
		Id   string
		Type string
	}
	MutationTimestamps struct {
		CreateTime time.Time
		UpdateTime time.Time
		DeleteTime time.Time
	}
	UserInCommonDetails struct {
		DisplayName string
		Email       string
	}
	UserHasLoggedInWithGithub struct {
		HasLoggedInWithGithub bool
	}
	UserHasLoggedInWithEmail struct {
		HasLoggedInWithEmail bool
	}
	UserHasLoggedInWithApple struct {
		HasLoggedInWithApple bool
	}
	UserView struct {
		IdentifyingMetadata
		MutationTimestamps
		UserInCommonDetails
		UserHasLoggedInWithGithub
		UserHasLoggedInWithEmail
		UserHasLoggedInWithApple
	}
	UserGithubDetails struct {
		githubToken string
	}
	UserEmailDetails struct {
		password string
	}
	UserAppleDetails struct {
		appleTokenAuthorizationCode string
	}
	UserLoginData struct {
		UserGithubDetails
		UserEmailDetails
		UserAppleDetails
	}
	UserCreate struct {
		UserInCommonDetails
		UserLoginData
	}
	UserLinkedAccounts struct {
		githubId string
		password string
		appleId  string
	}
	UserDocument struct {
		IdentifyingMetadata
		MutationTimestamps
		UserInCommonDetails
		UserLinkedAccounts
	}
)
