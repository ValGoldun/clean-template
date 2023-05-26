package di

import "github.com/ValGoldun/clean-template/pkg/clerk"

func ProvideClerk() clerk.Clerk {
	return clerk.New()
}
