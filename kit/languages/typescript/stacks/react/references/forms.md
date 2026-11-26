# React: forms and saving changes

For React projects that use MUI. Where this project's own code or
instructions differ, follow the project. These apply at every stability level.

## While saving

- Show progress on the button that started the save, with its loading state,
  and disable it until the save finishes, so it can't be submitted twice.
- Keep the form's values on screen while saving, and after a failure.

## Validation

- Check what the browser can for immediate feedback, such as required fields
  and formats, and show each problem on its field.
- The backend has the final say. When it rejects the input, show its errors
  on the matching fields where it says which field is wrong, and above the
  form otherwise. Never clear what the user typed.

## After saving

- Confirm success with a short `Snackbar`, or by going to the saved result.
- Refresh, or update in place, anything else on screen the change affects,
  such as the list the item appears in.
- Don't update the screen before the save succeeds (optimistic updates)
  unless the user asks for it.

## Destructive actions

Deleting, or anything else that can't be undone, asks first in a `Dialog`
that names what will be lost. The confirming button says what it does, such
as "Delete estimate", not "OK".
