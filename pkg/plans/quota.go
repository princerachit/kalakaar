package plans

// AccountLevelQuota is a struct that represents the account level quota
type AccountLevelQuota struct {
	// MaxConcurrentJobs is the maximum number of concurrent jobs that can be executed
	MaxConcurrentJobs int

	// MaxAllowedTimePerJob is the maximum allowed time for a job to execute
	MaxAllowedTimePerJob int // in seconds

	// UnsupportedAccountLevelQuotaSettings is the settings that are not supported yet
	UnsupportedAccountLevelQuotaSettings
}

type UnsupportedAccountLevelQuotaSettings struct {
	// maxConcurrentJobsPerUser is the maximum number of concurrent jobs that can be executed per user
	maxConcurrentJobsPerUser int // not supported yet
	// maxConcurrentJobsPerGroup is the maximum number of concurrent jobs that can be executed per group
	maxConcurrentJobsPerGroup int // not supported yet
	// maxConcurrentJobsPerUserPerGroup is the maximum number of concurrent jobs that can be executed per user per group
	maxConcurrentJobsPerUserPerGroup int // not supported yet

}

// IAccountLevelQuota is an interface that represents the account level quota
type IAccountLevelQuota interface {
	GetMaxConcurrentJobs() int
	GetMaxAllowedTimePerJob() int
}

type PipelineLevelQuota struct {
	// MaxCycles is the maximum number of cycles that can be executed
	MaxCycles int
	// MaxParallelPedals is the maximum number of pedals that can be executed in parallel
	MaxParallelPedals int
	// maxCPU is the maximum number of CPUs that can be used
	maxCPU int // not supported yet
}

type IPipelineLevelQuota interface {
	// GetMaxCycles is the maximum number of cycles that can be executed
	GetMaxCycles() int
	// GetMaxParallelPedals is the maximum number of pedals that can be executed in parallel
	GetMaxParallelPedals() int
	// getMaxCPU is the maximum number of CPUs that can be used
	getMaxCPU() int // not supported yet
}
