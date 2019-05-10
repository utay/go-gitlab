package gitlab

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
)

// ProjectImportExportService handles communication with the project import/export
// related methods of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ee/user/project/settings/import_export.html
type ProjectImportExportService struct {
	client *Client
}

// ExportProjectsOptions represents the available ExportProject() options.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_import_export.html#schedule-an-export
type ExportProjectsOptions struct {
	Description *string `url:"description,omitempty" json:"description,omitempty"`
	Upload      struct {
		Url        *string `json:"url,omitempty"`
		HttpMethod *string `json:"http_method,omitempty"`
	} `json:"upload,omitempty"`
}

func (s *ProjectImportExportService) ExportProject(pid interface{}, opt *ExportProjectsOptions, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/export", url.QueryEscape(project))
	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, err
}

// ProjectExportStatus represents a project export status.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_import_export.html#export-status
type ProjectExportStatus struct {
	ID                int     `json:"id"`
	Description       *string `json:"description,omitempty"`
	Name              *string `json:"name,omitempty"`
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`
	Path              *string `json:"path,omitempty"`
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`
	CreateAt          *string `json:"create_at,omitempty"`
	ExportStatus      *string `json:"export_status,omitempty"`
	Message           *string `json:"message,omitempty"`
	Links             struct {
		ApiURL *string `json:"api_url,omitempty"`
		WebURL *string `json:"web_url,omitempty"`
	} `json:"_links,omitempty"`
}

func (s ProjectExportStatus) String() string {
	return Stringify(s)
}

// GetProjectExportStatus Get the status of export.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_import_export.html#export-status
func (s *ProjectImportExportService) GetExportStatus(pid interface{}, options ...OptionFunc) (*ProjectExportStatus, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/export", url.QueryEscape(project))
	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}
	pb := new(ProjectExportStatus)
	resp, err := s.client.Do(req, pb)
	if err != nil {
		return nil, resp, err
	}
	return pb, resp, err
}

// GetProjectExportDownload Download the finished export.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_import_export.html#export-download
func (s *ProjectImportExportService) DownloadExport(pid interface{}, options ...OptionFunc) (io.Reader, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/export/download", url.QueryEscape(project))
	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}
	artifactsBuf := new(bytes.Buffer)
	resp, err := s.client.Do(req, artifactsBuf)
	if err != nil {
		return nil, resp, err
	}
	return artifactsBuf, resp, err
}

// ImportFileOptions represents the available GetProjectImportFile() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_import_export.html#import-a-file
type ImportFileOptions struct {
	Namespace      *string               `json:"namespace,omitempty"`
	File           *string               `json:"file"`
	Path           *string               `json:"path"`
	Overwrite      *bool                 `json:"overwrite,omitempty"`
	OverrideParams *CreateProjectOptions `json:"override_params, omitempty"`
}

func (s ImportFileOptions) String() string {
	return Stringify(s)
}

// ProjectImportStatus represents a project export status.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_import_export.html#import-status
type ProjectImportStatus struct {
	ID                int     `json:"id"`
	Description       *string `json:"description,omitempty"`
	Name              *string `json:"name,omitempty"`
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`
	Path              *string `json:"path,omitempty"`
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`
	CreateAt          *string `json:"create_at,omitempty"`
	ImportStatus      *string `json:"import_status,omitempty"`
}

func (s ProjectImportStatus) String() string {
	return Stringify(s)
}
// GetProjectImportFile import the project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_import_export.html#import-a-file
func (s *ProjectImportExportService) ImportProject(opt *ImportFileOptions, options ...OptionFunc) (*ProjectImportStatus, *Response, error) {
	req, err := s.client.NewRequest("POST", "/projects/import", opt, options)
	if err != nil {
		return nil, nil, err
	}
	pb := new(ProjectImportStatus)
	resp, err := s.client.Do(req, pb)
	if err != nil {
		return nil, resp, err
	}
	return pb, resp, err
}

// GetProjectImportStatus Get the status of an import.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/project_import_export.html#import-status
func (s *ProjectImportExportService) GetImportStatus(pid interface{}, options ...OptionFunc) (*ProjectImportStatus, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/import", url.QueryEscape(project))
	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}
	pb := new(ProjectImportStatus)
	resp, err := s.client.Do(req, pb)
	if err != nil {
		return nil, resp, err
	}
	return pb, resp, err
}
