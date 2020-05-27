package completeduploads

// Service represents the repository where uploaded objects are tracked
type Service struct {
	repo Repository
}

// NewService created a Service to track uploaded objects
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

// Close closes the service.
//
// No operation could be done after that.
func (s *Service) Close() error {
	return s.repo.Close()
}

// IsAlreadyUploaded checks if the file was already uploaded
func (s *Service) IsAlreadyUploaded(filePath string) (bool, error) {
	// find a previous upload in the repository
	_, err := s.repo.Get(filePath)
	if err != nil {
		// this file was not uploaded before
		return false, nil
	}

	return true, nil
}

// CacheAsAlreadyUploaded marks a file as already uploaded to prevent re-uploads
func (s *Service) CacheAsAlreadyUploaded(filePath string) error {
	item, err := NewCompletedUploadedFileItem(filePath)
	if err != nil {
		return err
	}
	return s.repo.Put(item)
}

// RemoveAsAlreadyUploaded removes a file previously marked as uploaded
func (s *Service) RemoveAsAlreadyUploaded(filePath string) error {
	return s.repo.Delete(filePath)
}
