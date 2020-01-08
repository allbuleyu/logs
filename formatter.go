package logs


type Formatter interface {
	Format() string
}
