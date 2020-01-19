package logs


type Formatter interface {
	Format(entry *Entry) ([]byte, error)
}
