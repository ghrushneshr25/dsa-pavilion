import React, {type ReactNode} from 'react';
import type {Props} from '@theme/Icon/Close';

export default function IconClose({
  width = 18,
  height = 18,
  className,
  ...restProps
}: Props): ReactNode {
  return (
    <svg
      className={className}
      width={width}
      height={height}
      viewBox="0 0 24 24"
      fill="none"
      aria-hidden="true"
      {...restProps}>
      <path
        d="M6 6L18 18M18 6L6 18"
        stroke="currentColor"
        strokeWidth="1.6"
      />
    </svg>
  );
}
