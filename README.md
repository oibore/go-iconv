# PACKAGE

    import "github.com/sloonz/go-iconv"

Bindings for iconv. Iconv is a set of functions used to convert strings
between different character sets

# FUNCTIONS

    func Conv(input string, tocode string, fromcode string) (string, error)

Utility function which create a codec, convert the string and then
destroy it


# TYPES

    type Iconv struct {
        // contains filtered or unexported fields
    }

Opaque structure containing the internal state of the codec

    func Open(tocode string, fromcode string) (*Iconv, error)

Create a codec which convert a string encoded in fromcode into a string
encoded in tocode

If you add //TRANSLIT at the end of tocode, any character which doesn't
exists in the destination charset will be replaced by its closest
equivalent (for example, € will be represented by EUR in ASCII). Else,
such a character will trigger an error.

    func (cd *Iconv) Close() error

Destroy the internal state of the codec, releasing associated memory

    func (cd *Iconv) Conv(input string) (result string, err error)

Use the codec to convert a string
