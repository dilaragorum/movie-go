generate-mocks:
	mockgen -source service/movie_service_interface.go -destination service/mock_movie_service.go -package service
	mockgen -source repository/movie_repository_interface.go -destination repository/mock_movie_repository.go -package repository