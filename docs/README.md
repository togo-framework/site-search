# search — docs

Bilingual Postgres full-text site search: to_tsvector('simple') + prefix tsquery for match-as-you-type + ILIKE fallback, ranked across registered content tables.

- **Source:** https://github.com/togo-framework/site-search
- **Reference app:** https://fadymondy.com

## Install

```bash
togo install togo-framework/site-search
```

The plugin self-registers on blank-import (no manual wiring). See the top-level
README for usage. This plugin was extracted from the fadymondy.com togo app as
part of turning that site into a reusable plugin suite.
