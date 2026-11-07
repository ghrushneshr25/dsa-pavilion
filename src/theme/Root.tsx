import React, {type ReactNode} from 'react';

export default function Root({children}: {children: ReactNode}): ReactNode {
  return (
    <>
      <div className="dsa-chrome" aria-hidden="true">
        <span>ctrl.plane v1.0</span>
        <span>dsa.pavilion</span>
        <span>sys.ok</span>
      </div>
      {children}
    </>
  );
}
