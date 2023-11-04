### What is this fork

This fork is a personal modded version of [gotosocial](https://github.com/superseriousbusiness/gotosocial).

This fork deploy to [@nyarla@kalaclista.com](https://kalaclista.com/@nyarla) as a personal fedivese instance.

### Modded features

- Imprpove federation compatibility for that doesn't support authorized fetch
  - But this feature make weaken to gotosocial privacy protection.

#### Configurations

- `kalaclista-allowed-unauthorized-get`
  - a boolean flag of turn off authorized fetch

### Development

This modded version uses `git rebase` for track to upstream.

Describe to branch means are:

- `main` - tracking branch for upstream gotosocial
- `kakaclista` - default and merged branch of my mods
- `unauthorized-fetch` - development branch for improve federation compatibility
- `readme` - documentation branch for my modded version
- `templates` - branch of make custom templates

### Maintainer

OKAMURA Naoki aka nyarla / [@nyarla@kalaclista.com](https://kalaclista.com/@nyarla)
