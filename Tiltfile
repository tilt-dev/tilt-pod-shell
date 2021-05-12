def trigger_on_ready(target, command, resource_name='ready_watcher'):
  watcher_dir = os.path.dirname(__file__)
  j = encode_json([{
    'target': target,
    'command': command,
  }])
  cmd = 'echo %s | (cd %s && go run ./cmd/tilt-trigger-on-ready)' % (shlex.quote(j), shlex.quote(watcher_dir))
  # TODO - figure out less clunky way to support multiple
  local_resource(resource_name, serve_cmd=cmd)
