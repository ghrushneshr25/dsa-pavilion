import React, {type ReactNode} from 'react';
import type {Props} from '@theme/Icon/Menu';

export default function IconMenu({
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
        d="M4 7H20M4 12H20M4 17H20"
        stroke="currentColor"
        strokeWidth="1.6"
      />
    </svg>
  );
}
