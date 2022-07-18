import React from 'react';

// Building blocks components main page
import IndexView from 'blocks/IndexView';

const routes = [
  {
    path: '/blocks',
    renderer: (params = {}): JSX.Element => <IndexView {...params} />,
  },
];

export default routes;
