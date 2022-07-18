import React from 'react';
import { Routes as ReactRoutes, Route, Navigate } from 'react-router-dom';
import viewsRoutes from 'views/routes';
import blocksRoutes from 'blocks/routes';

const Routes = (): JSX.Element => {
  return (
    <ReactRoutes>
      {viewsRoutes.map((item, i) => (
        <Route key={i} path={item.path} element={item.renderer()} />
      ))}

      {blocksRoutes.map((item, i) => (
        <Route key={i} path={item.path} element={item.renderer()} />
      ))}

      <Route path="*" element={<Navigate replace to="/not-found-cover" />} />
    </ReactRoutes>
  );
};

export default Routes;
