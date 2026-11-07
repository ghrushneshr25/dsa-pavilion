import React, {
  useCallback,
  useEffect,
  useMemo,
  useRef,
  useState,
  type ReactNode,
} from 'react';
import {createPortal} from 'react-dom';
import Link from '@docusaurus/Link';
import useIsBrowser from '@docusaurus/useIsBrowser';
import {useHistory} from '@docusaurus/router';
import {useBaseUrlUtils} from '@docusaurus/useBaseUrl';

import indexJson from '@site/src/data/dsaSearchIndex.json';

type SearchEntry = {
  id: string;
  title: string;
  type: 'concept' | 'problem' | string;
  category: string;
  tags: string[];
};

const catalog = indexJson as SearchEntry[];

function SearchIcon(): ReactNode {
  return (
    <svg width={16} height={16} viewBox="0 0 24 24" fill="none" aria-hidden>
      <circle cx="11" cy="11" r="6.25" stroke="currentColor" strokeWidth="1.6" />
      <path d="M16 16L21 21" stroke="currentColor" strokeWidth="1.6" strokeLinecap="square" />
    </svg>
  );
}

function isTypingTarget(el: EventTarget | null): boolean {
  if (!(el instanceof HTMLElement)) {
    return false;
  }
  const tag = el.tagName;
  return tag === 'INPUT' || tag === 'TEXTAREA' || tag === 'SELECT' || el.isContentEditable;
}

function scoreEntry(entry: SearchEntry, query: string): number {
  const q = query.trim().toLowerCase();
  if (!q) {
    return 1;
  }
  const title = entry.title.toLowerCase();
  const tags = entry.tags.join(' ').toLowerCase();
  const category = entry.category.toLowerCase();
  const type = entry.type.toLowerCase();
  if (title === q) {
    return 100;
  }
  if (title.startsWith(q)) {
    return 80;
  }
  if (title.includes(q)) {
    return 60;
  }
  const tokens = q.split(/\s+/).filter(Boolean);
  if (tokens.every((t) => title.includes(t) || tags.includes(t) || category.includes(t) || type.includes(t))) {
    return 50;
  }
  if (tags.includes(q)) {
    return 40;
  }
  if (type.startsWith(q) || type.includes(q)) {
    return 30;
  }
  if (category.includes(q)) {
    return 20;
  }
  return 0;
}

type Props = {
  mobile?: boolean;
  onClick?: () => void;
};

export default function SearchNavbarItem({mobile, onClick}: Props): ReactNode {
  const isBrowser = useIsBrowser();
  const [open, setOpen] = useState(false);

  const openSearch = useCallback(() => {
    setOpen(true);
    onClick?.();
  }, [onClick]);

  const closeSearch = useCallback(() => {
    setOpen(false);
  }, []);

  useEffect(() => {
    if (!isBrowser || mobile) {
      return undefined;
    }
    const onKey = (event: KeyboardEvent) => {
      const cmdK = (event.metaKey || event.ctrlKey) && event.key.toLowerCase() === 'k';
      const slash = event.key === '/' && !event.metaKey && !event.ctrlKey && !event.altKey;
      if (cmdK || (slash && !isTypingTarget(event.target))) {
        event.preventDefault();
        setOpen(true);
      }
    };
    window.addEventListener('keydown', onKey);
    return () => window.removeEventListener('keydown', onKey);
  }, [isBrowser, mobile]);

  const shortcut = isBrowser && /Mac|iPhone|iPad/.test(navigator.platform) ? '⌘K' : 'Ctrl+K';

  const button = (
    <button
      type="button"
      className="dsa-nav-btn dsa-search-toggle"
      aria-label="Search problems and concepts"
      title={`Search (${shortcut})`}
      onClick={openSearch}>
      <SearchIcon />
    </button>
  );

  if (mobile) {
    return (
      <li className="menu__list-item">
        <button type="button" className="menu__link" onClick={openSearch}>
          Search
        </button>
        {isBrowser && open
          ? createPortal(
              <SearchDialog onClose={closeSearch} />,
              document.body,
            )
          : null}
      </li>
    );
  }

  return (
    <>
      <div className="navbar__item">{button}</div>
      {isBrowser && open
        ? createPortal(<SearchDialog onClose={closeSearch} />, document.body)
        : null}
    </>
  );
}

function SearchDialog({onClose}: {onClose: () => void}): ReactNode {
  const history = useHistory();
  const {withBaseUrl} = useBaseUrlUtils();
  const inputRef = useRef<HTMLInputElement>(null);
  const [query, setQuery] = useState('');
  const [active, setActive] = useState(0);

  const results = useMemo(() => {
    const ranked = catalog
      .map((entry) => ({entry, score: scoreEntry(entry, query)}))
      .filter((row) => row.score > 0)
      .sort((a, b) => b.score - a.score || a.entry.title.localeCompare(b.entry.title));
    return ranked.slice(0, 12).map((row) => row.entry);
  }, [query]);

  useEffect(() => {
    inputRef.current?.focus();
  }, []);

  useEffect(() => {
    setActive(0);
  }, [query]);

  useEffect(() => {
    const prev = document.body.style.overflow;
    document.body.style.overflow = 'hidden';
    const onKey = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        event.preventDefault();
        onClose();
      }
    };
    window.addEventListener('keydown', onKey);
    return () => {
      document.body.style.overflow = prev;
      window.removeEventListener('keydown', onKey);
    };
  }, [onClose]);

  const go = useCallback(
    (entry: SearchEntry) => {
      history.push(withBaseUrl(`/${entry.id}`));
      onClose();
    },
    [history, onClose, withBaseUrl],
  );

  const onKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'ArrowDown') {
      event.preventDefault();
      setActive((i) => Math.min(i + 1, Math.max(results.length - 1, 0)));
    } else if (event.key === 'ArrowUp') {
      event.preventDefault();
      setActive((i) => Math.max(i - 1, 0));
    } else if (event.key === 'Enter' && results[active]) {
      event.preventDefault();
      go(results[active]);
    }
  };

  return (
    <div className="dsa-search" role="dialog" aria-modal="true" aria-label="Search">
      <button type="button" className="dsa-search__backdrop" aria-label="Close search" onClick={onClose} />
      <div className="dsa-search__panel">
        <div className="dsa-search__bar">
          <SearchIcon />
          <input
            ref={inputRef}
            className="dsa-search__input"
            value={query}
            onChange={(event) => setQuery(event.target.value)}
            onKeyDown={onKeyDown}
            placeholder="Search problems and concepts"
            aria-label="Search problems and concepts"
            autoComplete="off"
            spellCheck={false}
          />
          <kbd className="dsa-search__kbd">esc</kbd>
        </div>
        <ul className="dsa-search__list" role="listbox">
          {results.length === 0 ? (
            <li className="dsa-search__empty">No matching problem or concept.</li>
          ) : (
            results.map((entry, i) => (
              <li key={entry.id}>
                <SearchHit
                  entry={entry}
                  active={i === active}
                  onMouseEnter={() => setActive(i)}
                  onSelect={() => go(entry)}
                />
              </li>
            ))
          )}
        </ul>
      </div>
    </div>
  );
}

function SearchHit({
  entry,
  active,
  onMouseEnter,
  onSelect,
}: {
  entry: SearchEntry;
  active: boolean;
  onMouseEnter: () => void;
  onSelect: () => void;
}): ReactNode {
  return (
    <Link
      className={active ? 'dsa-search__hit is-active' : 'dsa-search__hit'}
      to={`/${entry.id}`}
      role="option"
      aria-selected={active}
      onMouseEnter={onMouseEnter}
      onClick={onSelect}>
      <span className="dsa-search__hit-type">{entry.type}</span>
      <span className="dsa-search__hit-body">
        <span className="dsa-search__hit-title">{entry.title}</span>
        <span className="dsa-search__hit-meta">
          {entry.category}
          {entry.tags.length ? ` · ${entry.tags.join(', ')}` : ''}
        </span>
      </span>
    </Link>
  );
}
