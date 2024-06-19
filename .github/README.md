### GoToSocial kalaclista modded edition

This is a fork of [gotosocial](https://github.com/superseriousbusiness/gotosocial) to my personal fediverse instance.

#### Modded features

- Improve federation compatibility by opt-out `AUTHORIZED_FETCH` mode.
  - this feature makes better federation between that doesn's support `AUTHORIZED_FETCH` by default.
  - but it makes weaken to gotosocial privacy guard.
- Keep remote emojis forever to media cache
  - this feature is just for me

##### Configurations

- `kalaclista-allowed-unauthorized-get` (default: false)
  - turn off to `AUTHORIZED_FETCH` compatible mode.
- `kalaclista-keep-emojis-forever` (default: false)
  - keep remote emojis to media cache forver

#### Quick hacks

- Skip to check Application Ids in post statuses.
  - for my instance, some DB value broken about Application ID by my mistake operation.

#### Development

This fork uses `git rebase` as development to support upstream,
and active branches are like as:

- `gotosocial-vX.Y.Z` - tracking branch to gotosocial
- `kalaclista-vX.Y.Z` - main branch of my fork
- `kalaclista-<topic>-vX.Y.Z` - this is the topic branch for my mods.

Rebase steps as:

```bash
# make base branch from upstream gotosocial
$ git switch main
$ git pull upstream main
$ git checkout -b gotosocial-vX.Y.Z {rev}

# make rebase branch from previous topic branch
$ git switch -c kalaclista-<topic>-vX.Y.Z kalaclista-<topic>-vX.Y.Z

# rebase new branch to upstream
$ git rebase gotosocial-vX.Y.Z

# testing...
$ VERSION=dev ./script/build.sh

# all tests ok, git push!
$ git push origin kalaclista-<topic>-vX.Y.Z

# finally, merge all topic branches to kalaclista branch
$ git switch -c kalaclista-vX.Y.Z gotosocial-vX.Y.Z
$ git rebase kalaclista-<topic>-vX.Y.Z # repeat it!
$ git push kalaclista-vX.Y.Z
```

#### Maintainer

OKAMURA Naoki aka nyarla / [@nyarla@kalaclista.com](https://kalaclista.com/@nyarla)
