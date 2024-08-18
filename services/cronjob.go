package services

import (
	"app/constants"
	"app/models"
	"app/repositories"
	"io"
	"net/http"
	"sync"

	"github.com/robfig/cron/v3"
)

type CronJobManagerService interface {
	Add(*Job) error
	Remove(*Job) error
	Update(*Job) error
}

func NewCronJobManager() CronJobManagerService {
	return &cronJobManager{
		Running: map[uint]chan bool{},
		Jobs:    map[uint]*Job{},
	}
}

type cronJobManager struct {
	mu      sync.Mutex
	Running map[uint]chan bool
	Jobs    map[uint]*Job
}

func (c *cronJobManager) Add(job *Job) error {
	exit := make(chan bool)
	c.mu.Lock()
	c.Jobs[job.ServiceID] = job
	go func() {
		c.Running[job.ServiceID] = exit
		// send HTTP Get Method  to check
		x := cron.New()
		// TODO: move config to environment
		// every minute
		x.AddFunc("* * * * *", func() {
			job.Check()
		})
		// Khởi động cron
		x.Start()
		<-exit
	}()
	c.mu.Unlock()
	return nil
}

func (c *cronJobManager) Remove(job *Job) error {
	c.Running[job.ServiceID] <- true
	c.mu.Lock()
	delete(c.Running, job.ServiceID)
	delete(c.Jobs, job.ServiceID)
	c.mu.Unlock()
	return nil
}

func (c *cronJobManager) Update(job *Job) error {
	c.Remove(job)
	c.Add(job)
	return nil
}

type Job struct {
	ServiceID   uint
	ServiceName string
	URL         string
	_repo       repositories.IServiceRepository
}

func (job *Job) SendHTTP() (*http.Response, string, error) {
	res, err := http.Get(job.URL)
	if err != nil {
		return res, "", err
	}

	defer res.Body.Close()
	body, errReadBody := io.ReadAll(res.Body)
	if errReadBody != nil {
		return nil, "", err
	}

	return res, string(body), nil
}

func (job *Job) Check() error {
	resp, body, err := job.SendHTTP()

	status := constants.Live
	if err != nil && resp.StatusCode != http.StatusOK {
		//log error
		status = constants.Dead
	}
	return job._repo.UpdateStatus(job.ServiceID, status, &models.LogChecked{
		HttpStatus:   http.StatusText(resp.StatusCode),
		ResponseTXT:  body,
		ServicesID:   int(job.ServiceID),
		Status:       int(status),
		RuntimeError: err.Error(),
	})
}
