import React, {type ReactNode} from 'react';
import type {Props} from '@theme/Icon/Copy';

export default function IconCopy(props: Props): ReactNode {
  return (
    <svg viewBox="0 0 24 24" fill="none" aria-hidden="true" {...props}>
      <rect x="8" y="8" width="12" height="12" stroke="currentColor" strokeWidth="1.6" />
      <path d="M4 16V4H16" stroke="currentColor" strokeWidth="1.6" />
    </svg>
  );
}
