package configuration

// Settings is a struct that represents the settings for the repository
type Settings struct {
	// MaxConcurrentJobs is the maximum number of concurrent jobs that can be executed
	MaxConcurrentJobs int

	// Below are not supported yet and not guaranteed to be supported in future

	// MaxConcurrentJobsPerUser is the maximum number of concurrent jobs that can be executed per user
	MaxConcurrentJobsPerUser int // not supported yet
	// MaxConcurrentJobsPerGroup is the maximum number of concurrent jobs that can be executed per group
	MaxConcurrentJobsPerGroup int // not supported yet
	// MaxConcurrentJobsPerUserPerGroup is the maximum number of concurrent jobs that can be executed per user per group
	MaxConcurrentJobsPerUserPerGroup int // not supported yet

	// MaxAllowedTimePerJob is the maximum allowed time for a job to execute
	MaxAllowedTimePerJob int // in seconds
}
