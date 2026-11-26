# React: calling services

For React projects, adding to the `typescript` skill's
[services.md](../../typescript/references/services.md). Where this project's
own code or instructions differ, follow the project.

Pages and components call the service functions. They never call `fetch`,
build a service's URLs, read status codes, or call an SDK directly. The
exception is an SDK's own React components and hooks, such as Clerk's, which
pages use as documented.

Pass each request's `AbortSignal` from the component that starts it (see
[data.md](data.md)), and describe a list in the app's own terms (see
[lists.md](lists.md)).
