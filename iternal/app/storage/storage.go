package storage

type Storage interface {
	News() NewsRepository
	Users() UsersRepository
	Calendar() CalendarRepository
}
