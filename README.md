# ArooToSocial - Source

Arooooo!!

![Confused wolf gif](https://github.com/jwbjnwolf/arootosocial/assets/59905825/f763c01e-0f7c-4dfd-82c9-44bb7ec649fd)

- ArooToSocial: https://github.com/jwbjnwolf/arootosocial,
- GoToSocial: https://github.com/superseriousbusiness/gotosocial.

# Changes:
- Make unlisted posts viewable via link in the web interface, (still hidden from profile view): [27ba6ff7987c3b7cb3823d677afff8a716e928d7](https://github.com/jwbjnwolf/arootosocial-src/commit/27ba6ff7987c3b7cb3823d677afff8a716e928d7).
- Make local only posts viewable via the web interface, (still hides from profile view even if public unless pinned): [a4b162fcb059f23bca638eb25f6fc695e319b885](https://github.com/jwbjnwolf/arootosocial-src/commit/a4b162fcb059f23bca638eb25f6fc695e319b885).
- Change animated emojis to only animate on hover/tap in the web interface, whilst still respecting `prefers-reduced-motion`: [5a9095d49a0363fe120dada4ea0ba1a499c2065b](https://github.com/jwbjnwolf/arootosocial-src/commit/5a9095d49a0363fe120dada4ea0ba1a499c2065b/).

# Building:
- Clone repo locally along with upstream added as remote "upstream".
- Install Go.
- Change working directory in terminal to the cloned repo.
- Build the binary file for a stated platform & arch: `GOOS=linux GOARCH=amd64 ./scripts/build.sh`.
- I've added some convience additional scripts `build_amd64_linux.sh` & `build_arm64_linux.sh` though so you can just run those instead of the `build.sh` directly with those variables infront.
- To ensure you have the correct binary generated, use `file` command: `file gotosocial`.