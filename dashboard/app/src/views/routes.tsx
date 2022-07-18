import React from 'react';

import {
  IndexView,
} from 'views';

const routes = [
  {
    path: '/',
    renderer: (params = {}): JSX.Element => <IndexView {...params} />,
  },
];

export default routes;
