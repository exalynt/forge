# Keep the focus off a browser that's opening under WSLg. chrome.sh runs this on
# the Windows side, where every WSLg window belongs to msrdc.exe and takes the
# foreground as it opens. Chrome can't be told not to, so while the browser
# opens, focus goes straight back to the window the user was last in.
#
# Prints "ready" once it's watching, so Chrome can wait for that. Windows open
# within a second or so, and it watches for longer in case the machine is busy.
# Until it stops, clicking into a browser window also gets sent back.
$ProgressPreference = 'SilentlyContinue'
Add-Type -Namespace Forge -Name Win32 -MemberDefinition @'
[DllImport("user32.dll")] public static extern IntPtr GetForegroundWindow();
[DllImport("user32.dll")] public static extern bool SetForegroundWindow(IntPtr window);
[DllImport("user32.dll")] public static extern uint GetWindowThreadProcessId(IntPtr window, out uint process);
'@

function Test-Msrdc($window) {
  $id = [uint32]0
  [void][Forge.Win32]::GetWindowThreadProcessId($window, [ref]$id)
  (Get-Process -Id $id -ErrorAction SilentlyContinue).ProcessName -eq 'msrdc'
}

$keep = [Forge.Win32]::GetForegroundWindow()
'ready'
$timer = [Diagnostics.Stopwatch]::StartNew()
while ($timer.Elapsed.TotalSeconds -lt 10) {
  $now = [Forge.Win32]::GetForegroundWindow()
  if ($now -ne $keep -and $now -ne [IntPtr]::Zero) {
    if (Test-Msrdc $now) {
      [void][Forge.Win32]::SetForegroundWindow($keep)
    } else {
      # The user moved on to another window, so that's the one to go back to.
      $keep = $now
    }
  }
  Start-Sleep -Milliseconds 30
}
