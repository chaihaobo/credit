// Package repository
// @author： Boice
// @createTime：2022/11/28 11:20
package repository

import "credit-platform/resource"

type (
	Repository interface {
		Client
		Customer() CustomerRepository
		CustomerApi() CustomerApiRepository
		ApiCallLog() ApiCallLogRepository
		Api() ApiRepository
	}
	repository struct {
		Client
		customer    CustomerRepository
		customerApi CustomerApiRepository
		api         ApiRepository
		apiCallLog  ApiCallLogRepository
	}
)

func (r *repository) ApiCallLog() ApiCallLogRepository {
	return r.apiCallLog
}

func (r *repository) Api() ApiRepository {
	return r.api
}

func (r *repository) CustomerApi() CustomerApiRepository {
	return r.customerApi
}

func (r *repository) Customer() CustomerRepository {
	return r.customer
}

func New(res resource.Resource) Repository {
	cli := NewClient(res.Config())
	return &repository{
		Client:      cli,
		customer:    newCustomerRepository(res, cli),
		customerApi: newCustomerApiRepository(res, cli),
		api:         newApiRepository(res, cli),
		apiCallLog:  newApiCallLogRepository(res, cli),
	}
}
