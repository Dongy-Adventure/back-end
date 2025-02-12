# Server & DB for SE_II

## Table of Contents

- [Installation](#installation)
- [Contributing](#contributing)

## Installation

1. clone the repository:
   ```bash
   git clone git@github.com:Dongy-s-Advanture/back-end.git
   ```
2. install docker & go if you haven't
3. set up .env file & init.js
4. run docker
   ```bash
   docker-compose up --build -d
   ```
5. run server
   ```bash
   go run .\cmd\main.go\
   # or air if you have installed
   air
   ```

## Contributing

Commit messages should follow one of the following types:

- feat: A new feature.
- fix: A bug fix.
- refactor: A code change that neither fixes a bug nor adds a feature.
- style: Changes that do not affect the meaning of the code (e.g., whitespace, formatting).
- docs: Documentation only changes.
- chore: Changes to the build process or auxiliary tools and libraries.
- Example commit messages:

```bash
git commit -m "feat: <what-you-did>"
# or
git commit -m "fix: <what-you-fixed>"
# or
git commit -m "refactor: <what-you-refactored>"
```

For more details, refer to the [Conventional commit](https://www.conventionalcommits.org/en/v1.0.0/) format documentation.
