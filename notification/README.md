# Estuary Metrics Notification API

Simple API that sends notifications to the user to registered websocket subscribers.

This is a standalone api module that runs on its own daemon.

## Postgres Trigger
In order to use this module, you need to create a postgres trigger and the LISTEN/NOTIFY functions are enabled.