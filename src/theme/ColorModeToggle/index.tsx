import React, {useCallback, useEffect, useRef, useState, type ReactNode} from 'react';
import clsx from 'clsx';
import useIsBrowser from '@docusaurus/useIsBrowser';
import {useColorMode} from '@docusaurus/theme-common';
import type {Props} from '@theme/ColorModeToggle';
import type {ColorMode} from '@docusaurus/theme-common';

const TRANSIT_MS = 1400;
const THEME_SWAP_AT = 0.12;
const FADE_OUT_MS = 480;

function SunIcon(): ReactNode {
  return (
    <svg width={16} height={16} viewBox="0 0 24 24" fill="none" aria-hidden>
      <circle cx="12" cy="12" r="4" stroke="currentColor" strokeWidth="1.6" />
      <path
        d="M12 2V4M12 20V22M4 12H2M22 12H20M5.64 5.64L4.22 4.22M19.78 19.78L18.36 18.36M5.64 18.36L4.22 19.78M19.78 4.22L18.36 5.64"
        stroke="currentColor"
        strokeWidth="1.6"
        strokeLinecap="square"
      />
    </svg>
  );
}

function MoonIcon(): ReactNode {
  return (
    <svg width={16} height={16} viewBox="0 0 24 24" fill="none" aria-hidden>
      <path
        d="M20 14.5A8.5 8.5 0 1 1 9.5 4 7 7 0 0 0 20 14.5Z"
        stroke="currentColor"
        strokeWidth="1.6"
        strokeLinejoin="round"
      />
    </svg>
  );
}

function prefersReducedMotion(): boolean {
  return window.matchMedia('(prefers-reduced-motion: reduce)').matches;
}

function applyThemeMeta(theme: ColorMode) {
  const meta = document.querySelector('meta[name="theme-color"]');
  if (meta) {
    meta.setAttribute('content', theme === 'light' ? '#ffffff' : '#000000');
  }
}

function playThemeWave(next: ColorMode, onMid: () => void, onDone: () => void): () => void {
  if (prefersReducedMotion()) {
    onMid();
    onDone();
    return () => undefined;
  }

  const direction = next === 'light' ? 'to-light' : 'to-dark';
  const root = document.createElement('div');
  root.className = `theme-wave theme-wave--${direction}`;
  root.setAttribute('aria-hidden', 'true');
  root.innerHTML = '<div class="theme-wave__veil"></div>';

  document.body.appendChild(root);
  document.documentElement.classList.add('is-theme-waving');

  void root.offsetWidth;
  root.classList.add('is-active');

  const midTimer = window.setTimeout(onMid, TRANSIT_MS * THEME_SWAP_AT);
  const endTimer = window.setTimeout(() => {
    root.classList.add('is-done');
    window.setTimeout(() => {
      document.documentElement.classList.remove('is-theme-waving');
      root.remove();
      onDone();
    }, FADE_OUT_MS);
  }, TRANSIT_MS);

  return () => {
    window.clearTimeout(midTimer);
    window.clearTimeout(endTimer);
    document.documentElement.classList.remove('is-theme-waving');
    root.remove();
  };
}

function ColorModeToggle({buttonClassName}: Props): ReactNode {
  const isBrowser = useIsBrowser();
  const {colorMode, setColorMode} = useColorMode();
  const theme = colorMode === 'light' ? 'light' : 'dark';
  const [busy, setBusy] = useState(false);
  const themeRef = useRef<ColorMode>(theme);
  const cleanupRef = useRef<(() => void) | null>(null);

  useEffect(() => {
    themeRef.current = theme;
  }, [theme]);

  useEffect(
    () => () => {
      cleanupRef.current?.();
    },
    [],
  );

  const toggle = useCallback(() => {
    if (!isBrowser || busy) {
      return;
    }

    const next: ColorMode = themeRef.current === 'dark' ? 'light' : 'dark';

    setBusy(true);
    cleanupRef.current?.();
    cleanupRef.current = playThemeWave(
      next,
      () => {
        setColorMode(next);
        applyThemeMeta(next);
      },
      () => {
        setBusy(false);
        cleanupRef.current = null;
      },
    );
  }, [busy, isBrowser, setColorMode]);

  return (
    <div className="dsa-color-toggle">
      <button
        className={clsx('dsa-nav-btn', buttonClassName)}
        type="button"
        onClick={toggle}
        disabled={!isBrowser || busy}
        title={theme === 'dark' ? 'Switch to light mode' : 'Switch to dark mode'}
        aria-label={theme === 'dark' ? 'Switch to light mode' : 'Switch to dark mode'}
        aria-busy={busy}
      >
        {theme === 'dark' ? <SunIcon /> : <MoonIcon />}
      </button>
    </div>
  );
}

export default React.memo(ColorModeToggle);
