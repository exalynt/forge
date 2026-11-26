# React: search, filters, sorting, and pagination

For React projects that use MUI. Where this project's own code or
instructions differ, follow the project. [data.md](data.md) covers loading,
empty, and failed states, and cancelling stale requests, which lists need too.

## Let the API do the work

- When the API can search, filter, sort, or paginate a list, use it for that.
  Downloading everything to filter it in the browser gets slower as the data
  grows.
- Never sort or filter one page of paginated results in the browser. That only
  orders the page, not the list, and looks right until there's a second page.
- Filtering, sorting, or searching in the browser is fine for small, bounded
  sets, such as a fixed list of options or the handful of items one user has.
- If a list can grow and the API can't do what the page needs, tell the user.
  That's a change to the API, not something to work around in the browser.
- The page describes what it wants in the app's own terms: a search, filters,
  a sort, a page. The service function turns that into whatever the API
  takes, so pages never know the API's parameter names or paging scheme.

## Keep list state in the URL

- The search, filters, sort, and page live in the URL's query string, under
  the app's own names. Reloading, sharing the link, or coming back from a row
  shows the same list.
- Leave a value out of the URL when it's at its default, and fall back to the
  default when a value in the URL isn't valid.
- Replace the history entry when the list state changes, rather than adding
  one, so Back leaves the list instead of undoing each keystroke.
- Changing the search, a filter, the sort, or the page size goes back to the
  first page.

## Debounce typed input

- A search box updates on every keystroke. The URL and the request wait until
  typing pauses for about 300 ms.
- Selects, checkboxes, sort headers, and page buttons apply straight away,
  with no debounce.
- Trim the search, and treat an empty one as no search.
- Clear a pending debounce timer when the component unmounts.
- Each new query aborts the request before it (see [data.md](data.md)), and
  the list shows it's refreshing rather than blanking.

## Show what's applied

- When any filter or search is in use, show it, and offer a way to clear them
  all.
- Show the count of results when the API gives one.
