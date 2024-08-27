// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddScheduleActionRecord provides a mock function with given fields: ctx, scheduleActionRecord
func (_m *DBClient) AddScheduleActionRecord(ctx context.Context, scheduleActionRecord models.ScheduleActionRecord) (models.ScheduleActionRecord, errors.EdgeX) {
	ret := _m.Called(ctx, scheduleActionRecord)

	if len(ret) == 0 {
		panic("no return value specified for AddScheduleActionRecord")
	}

	var r0 models.ScheduleActionRecord
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, models.ScheduleActionRecord) (models.ScheduleActionRecord, errors.EdgeX)); ok {
		return rf(ctx, scheduleActionRecord)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ScheduleActionRecord) models.ScheduleActionRecord); ok {
		r0 = rf(ctx, scheduleActionRecord)
	} else {
		r0 = ret.Get(0).(models.ScheduleActionRecord)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ScheduleActionRecord) errors.EdgeX); ok {
		r1 = rf(ctx, scheduleActionRecord)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddScheduleActionRecords provides a mock function with given fields: ctx, scheduleActionRecord
func (_m *DBClient) AddScheduleActionRecords(ctx context.Context, scheduleActionRecord []models.ScheduleActionRecord) ([]models.ScheduleActionRecord, errors.EdgeX) {
	ret := _m.Called(ctx, scheduleActionRecord)

	if len(ret) == 0 {
		panic("no return value specified for AddScheduleActionRecords")
	}

	var r0 []models.ScheduleActionRecord
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []models.ScheduleActionRecord) ([]models.ScheduleActionRecord, errors.EdgeX)); ok {
		return rf(ctx, scheduleActionRecord)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []models.ScheduleActionRecord) []models.ScheduleActionRecord); ok {
		r0 = rf(ctx, scheduleActionRecord)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScheduleActionRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []models.ScheduleActionRecord) errors.EdgeX); ok {
		r1 = rf(ctx, scheduleActionRecord)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddScheduleJob provides a mock function with given fields: ctx, scheduleJob
func (_m *DBClient) AddScheduleJob(ctx context.Context, scheduleJob models.ScheduleJob) (models.ScheduleJob, errors.EdgeX) {
	ret := _m.Called(ctx, scheduleJob)

	if len(ret) == 0 {
		panic("no return value specified for AddScheduleJob")
	}

	var r0 models.ScheduleJob
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, models.ScheduleJob) (models.ScheduleJob, errors.EdgeX)); ok {
		return rf(ctx, scheduleJob)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ScheduleJob) models.ScheduleJob); ok {
		r0 = rf(ctx, scheduleJob)
	} else {
		r0 = ret.Get(0).(models.ScheduleJob)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ScheduleJob) errors.EdgeX); ok {
		r1 = rf(ctx, scheduleJob)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllScheduleActionRecords provides a mock function with given fields: ctx, start, end, offset, limit
func (_m *DBClient) AllScheduleActionRecords(ctx context.Context, start int64, end int64, offset int, limit int) ([]models.ScheduleActionRecord, errors.EdgeX) {
	ret := _m.Called(ctx, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for AllScheduleActionRecords")
	}

	var r0 []models.ScheduleActionRecord
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int, int) ([]models.ScheduleActionRecord, errors.EdgeX)); ok {
		return rf(ctx, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int, int) []models.ScheduleActionRecord); ok {
		r0 = rf(ctx, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScheduleActionRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllScheduleJobs provides a mock function with given fields: ctx, labels, offset, limit
func (_m *DBClient) AllScheduleJobs(ctx context.Context, labels []string, offset int, limit int) ([]models.ScheduleJob, errors.EdgeX) {
	ret := _m.Called(ctx, labels, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for AllScheduleJobs")
	}

	var r0 []models.ScheduleJob
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) ([]models.ScheduleJob, errors.EdgeX)); ok {
		return rf(ctx, labels, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) []models.ScheduleJob); ok {
		r0 = rf(ctx, labels, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScheduleJob)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, labels, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// DeleteScheduleActionRecordByAge provides a mock function with given fields: ctx, age
func (_m *DBClient) DeleteScheduleActionRecordByAge(ctx context.Context, age int64) errors.EdgeX {
	ret := _m.Called(ctx, age)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScheduleActionRecordByAge")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, int64) errors.EdgeX); ok {
		r0 = rf(ctx, age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// DeleteScheduleJobByName provides a mock function with given fields: ctx, name
func (_m *DBClient) DeleteScheduleJobByName(ctx context.Context, name string) errors.EdgeX {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScheduleJobByName")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) errors.EdgeX); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// LatestScheduleActionRecordsByJobName provides a mock function with given fields: ctx, jobName
func (_m *DBClient) LatestScheduleActionRecordsByJobName(ctx context.Context, jobName string) ([]models.ScheduleActionRecord, errors.EdgeX) {
	ret := _m.Called(ctx, jobName)

	if len(ret) == 0 {
		panic("no return value specified for LatestScheduleActionRecordsByJobName")
	}

	var r0 []models.ScheduleActionRecord
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.ScheduleActionRecord, errors.EdgeX)); ok {
		return rf(ctx, jobName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.ScheduleActionRecord); ok {
		r0 = rf(ctx, jobName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScheduleActionRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, jobName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleActionRecordCountByJobName provides a mock function with given fields: ctx, jobName
func (_m *DBClient) ScheduleActionRecordCountByJobName(ctx context.Context, jobName string) (uint32, errors.EdgeX) {
	ret := _m.Called(ctx, jobName)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleActionRecordCountByJobName")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint32, errors.EdgeX)); ok {
		return rf(ctx, jobName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint32); ok {
		r0 = rf(ctx, jobName)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, jobName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleActionRecordCountByJobNameAndStatus provides a mock function with given fields: ctx, jobName, status
func (_m *DBClient) ScheduleActionRecordCountByJobNameAndStatus(ctx context.Context, jobName string, status string) (uint32, errors.EdgeX) {
	ret := _m.Called(ctx, jobName, status)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleActionRecordCountByJobNameAndStatus")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (uint32, errors.EdgeX)); ok {
		return rf(ctx, jobName, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) uint32); ok {
		r0 = rf(ctx, jobName, status)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) errors.EdgeX); ok {
		r1 = rf(ctx, jobName, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleActionRecordCountByStatus provides a mock function with given fields: ctx, status
func (_m *DBClient) ScheduleActionRecordCountByStatus(ctx context.Context, status string) (uint32, errors.EdgeX) {
	ret := _m.Called(ctx, status)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleActionRecordCountByStatus")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint32, errors.EdgeX)); ok {
		return rf(ctx, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint32); ok {
		r0 = rf(ctx, status)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleActionRecordTotalCount provides a mock function with given fields: ctx
func (_m *DBClient) ScheduleActionRecordTotalCount(ctx context.Context) (uint32, errors.EdgeX) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleActionRecordTotalCount")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context) (uint32, errors.EdgeX)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint32); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleActionRecordsByJobName provides a mock function with given fields: ctx, jobName, start, end, offset, limit
func (_m *DBClient) ScheduleActionRecordsByJobName(ctx context.Context, jobName string, start int64, end int64, offset int, limit int) ([]models.ScheduleActionRecord, errors.EdgeX) {
	ret := _m.Called(ctx, jobName, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleActionRecordsByJobName")
	}

	var r0 []models.ScheduleActionRecord
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64, int, int) ([]models.ScheduleActionRecord, errors.EdgeX)); ok {
		return rf(ctx, jobName, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64, int, int) []models.ScheduleActionRecord); ok {
		r0 = rf(ctx, jobName, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScheduleActionRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, jobName, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleActionRecordsByJobNameAndStatus provides a mock function with given fields: ctx, jobName, status, start, end, offset, limit
func (_m *DBClient) ScheduleActionRecordsByJobNameAndStatus(ctx context.Context, jobName string, status string, start int64, end int64, offset int, limit int) ([]models.ScheduleActionRecord, errors.EdgeX) {
	ret := _m.Called(ctx, jobName, status, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleActionRecordsByJobNameAndStatus")
	}

	var r0 []models.ScheduleActionRecord
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, int64, int, int) ([]models.ScheduleActionRecord, errors.EdgeX)); ok {
		return rf(ctx, jobName, status, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64, int64, int, int) []models.ScheduleActionRecord); ok {
		r0 = rf(ctx, jobName, status, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScheduleActionRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, jobName, status, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleActionRecordsByStatus provides a mock function with given fields: ctx, status, start, end, offset, limit
func (_m *DBClient) ScheduleActionRecordsByStatus(ctx context.Context, status string, start int64, end int64, offset int, limit int) ([]models.ScheduleActionRecord, errors.EdgeX) {
	ret := _m.Called(ctx, status, start, end, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleActionRecordsByStatus")
	}

	var r0 []models.ScheduleActionRecord
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64, int, int) ([]models.ScheduleActionRecord, errors.EdgeX)); ok {
		return rf(ctx, status, start, end, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64, int, int) []models.ScheduleActionRecord); ok {
		r0 = rf(ctx, status, start, end, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ScheduleActionRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, status, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleJobById provides a mock function with given fields: ctx, id
func (_m *DBClient) ScheduleJobById(ctx context.Context, id string) (models.ScheduleJob, errors.EdgeX) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleJobById")
	}

	var r0 models.ScheduleJob
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.ScheduleJob, errors.EdgeX)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.ScheduleJob); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.ScheduleJob)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleJobByName provides a mock function with given fields: ctx, name
func (_m *DBClient) ScheduleJobByName(ctx context.Context, name string) (models.ScheduleJob, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleJobByName")
	}

	var r0 models.ScheduleJob
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.ScheduleJob, errors.EdgeX)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.ScheduleJob); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(models.ScheduleJob)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ScheduleJobTotalCount provides a mock function with given fields: ctx, labels
func (_m *DBClient) ScheduleJobTotalCount(ctx context.Context, labels []string) (uint32, errors.EdgeX) {
	ret := _m.Called(ctx, labels)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleJobTotalCount")
	}

	var r0 uint32
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, []string) (uint32, errors.EdgeX)); ok {
		return rf(ctx, labels)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) uint32); ok {
		r0 = rf(ctx, labels)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) errors.EdgeX); ok {
		r1 = rf(ctx, labels)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateScheduleJob provides a mock function with given fields: ctx, scheduleJob
func (_m *DBClient) UpdateScheduleJob(ctx context.Context, scheduleJob models.ScheduleJob) errors.EdgeX {
	ret := _m.Called(ctx, scheduleJob)

	if len(ret) == 0 {
		panic("no return value specified for UpdateScheduleJob")
	}

	var r0 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, models.ScheduleJob) errors.EdgeX); ok {
		r0 = rf(ctx, scheduleJob)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.EdgeX)
		}
	}

	return r0
}

// NewDBClient creates a new instance of DBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBClient {
	mock := &DBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
