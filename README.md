# ArooToSocial - Source

Arooooo!!

![Confused wolf gif](https://github.com/jwbjnwolf/arootosocial/assets/59905825/f763c01e-0f7c-4dfd-82c9-44bb7ec649fd)

- ArooToSocial: https://github.com/jwbjnwolf/arootosocial,
- GoToSocial: https://github.com/superseriousbusiness/gotosocial.

# Changes:
- Make unlisted posts viewable via link in the web interface, (still hidden from profile view): [bbcdade6daabfb17d5daae941494d111f9e041c3](https://github.com/jwbjnwolf/arootosocial-src/commit/bbcdade6daabfb17d5daae941494d111f9e041c3)

# Building:
- Clone repo locally along with upstream added as remote "upstream".
- Install Go.
- Change working directory in terminal to the cloned repo.
- Build the binary file for a stated platform & arch: `GOOS=linux GOARCH=amd64 ./scripts/build.sh`
