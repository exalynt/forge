# React: loading and showing data

For React projects that use MUI. Where this project's own code or
instructions differ, follow the project. These apply at every stability level.

Data comes from the services layer (see [services.md](services.md)). Nothing here
assumes a particular backend.

## Every view has four states

Anything that shows fetched data handles all four: loading, loaded, empty, and
failed.

- **Loading, the first time:** a `Skeleton` shaped like the content it stands
  in for, so the layout doesn't jump when the data arrives. Use a spinner only
  where there's no content shape to stand in for, such as a button or a
  dialog's whole body.
- **Loading again** (a new search, sort, page, or retry): keep what's on
  screen and mark it as refreshing, such as a `LinearProgress` above it or
  dimmed rows. Never blank it back to a skeleton or spinner.
- **Empty:** say why there's nothing and what to do next, such as a button to
  create the first one. "No results for these filters" is a different state
  from "nothing yet", with a way to clear the filters.
- **Failed:** an `Alert` saying what couldn't load, with a retry. Never a
  blank area or a view stuck loading. Don't show raw error text from the
  backend unless it's written for the user.

## No flashing

Wait about 200 ms before showing any loading indicator, skeleton or spinner,
so a fast response goes straight to the content without a flash. Until then,
keep the space the content will take, so nothing shifts.

## Stale responses never win

- Every request gets an `AbortController`, and its signal is passed to the
  service function. Abort it when the component unmounts or when whatever the
  request depends on changes, such as the search or the record's ID.
- Ignore an aborted request's result and its error. Don't set state from it,
  and don't report it as a failure.
- A response for an old query must never replace the results of a newer one.

## Pages with several kinds of data

A page that shows several independent things, such as a record, its
activity, and some related totals, should show each as soon as it's ready,
not all of them once the slowest has loaded.

- Render the page's layout straight away. Each independent part fetches its
  own data and handles its own four states, so it fills in when its data
  arrives.
- Start independent requests at the same time. Avoid waterfalls: a request
  that doesn't need another's result never waits for it, whether by awaiting
  them one after the other or by a child that only mounts, and only then
  starts fetching, once its parent's data has loaded. If a part only needs an
  ID the page already has, such as one from the route, it fetches right away.
- Chain requests only when one really needs another's result.
- Data one part needs all of before it can show anything is fetched together,
  with `Promise.all`, behind one state.
- Parts that aren't visible yet, such as a closed tab, a collapsed panel, or a
  dialog, fetch when they're opened, not with the page.
- A part that fails shows its own error and retry, and the rest of the page
  keeps working. Only the data the page can't exist without, such as the
  record a `View` page is about, fails the whole page, and a record that
  doesn't exist shows a not-found state.
- When two parts need the same data, fetch it once in their nearest shared
  parent and pass it down. Don't add a cache for it.
- Each part's skeleton is sized like its content, so parts that load later
  don't push others around.
