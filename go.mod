module github.com/willoong9559/Gotunnel

go 1.17

replace github.com/willoong9559/Gotunnel/tunnel => ./tunnel

replace github.com/willoong9559/Gotunnel/utils/aead => ./utils/aead

require golang.org/x/crypto v0.0.0-20220214200702-86341886e292

require golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
