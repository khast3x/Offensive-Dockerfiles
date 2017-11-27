# Golismero

## Source
https://github.com/golismero/golismero

## Usage

```bash
cd golismero/
docker build -t golismero .
docker run -it golismero:latest
```

## Help

```bash
usage: golismero.py
       [-h]
       [--help]
       [-f FILE]
       [--config FILE]
       [--user-config FILE]
       [-p NAME]
       [--ui-mode MODE]
       [-v]
       [-q]
       [--color]
       [--no-color]
       [--audit-name NAME]
       [-db DATABASE]
       [-nd]
       [-i FILENAME]
       [-ni]
       [-o FILENAME]
       [-no]
       [--full]
       [--brief]
       [--allow-subdomains]
       [--forbid-subdomains]
       [--parent]
       [-np]
       [-r DEPTH]
       [--follow-redirects]
       [--no-follow-redirects]
       [--follow-first]
       [--no-follow-first]
       [--max-connections MAX_CONNECTIONS]
       [-l MAX_LINKS]
       [-pu USER]
       [-pp PASS]
       [-pa ADDRESS]
       [-pn PORT]
       [--cookie COOKIE]
       [--user-agent USER_AGENT]
       [--cookie-file FILE]
       [--persistent-cache]
       [--volatile-cache]
       [-a PLUGIN:KEY=VALUE]
       [-e PLUGIN]
       [-d PLUGIN]
       [--max-concurrent N]
       [--plugin-timeout N]
       [--plugins-folder PATH]
       COMMAND
       [TARGET [TARGET ...]]

available commands:

  SCAN:
    Perform a vulnerability scan on the given targets. Optionally import
    results from other tools and write a report. The arguments that follow may
    be domain names, IP addresses or web pages.

  RESCAN:
    Same as SCAN, but previously run tests are repeated. If the database is
    new, this command is identical to SCAN.

  PROFILES:
    Show a list of available config profiles. This command takes no arguments.

  PLUGINS:
    Show a list of available plugins. This command takes no arguments.

  INFO:
    Show detailed information on a given plugin. The arguments that follow are
    the plugin IDs. You can use glob-style wildcards.

  REPORT:
    Write a report from an earlier scan. This command takes no arguments.
    To specify output files use the -o switch.

  IMPORT:
    Import results from other tools and optionally write a report, but don't
    scan the targets. This command takes no arguments. To specify input files
    use the -i switch.

  DUMP:
    Dump the database from an earlier scan in SQL format. This command takes no
    arguments. To specify output files use the -o switch.

  LOAD:
    Load a database dump from an earlier scan in SQL format. This command takes
    no arguments. To specify input files use the -i switch.

  UPDATE:
    Update GoLismero to the latest version. Requires Git to be installed and
    available in the PATH. This command takes no arguments.

examples:

  scan a website and show the results on screen:
    golismero.py scan http://www.example.com

  grab Nmap results, scan all hosts found and write an HTML report:
    golismero.py scan -i nmap_output.xml -o report.html

  grab results from OpenVAS and show them on screen, but don't scan anything:
    golismero.py import -i openvas_output.xml

  show a list of all available configuration profiles:
    golismero.py profiles

  show a list of all available plugins:
    golismero.py plugins

  show information on all bruteforcer plugins:
    golismero.py info brute_*

  dump the database from a previous scan:
    golismero.py dump -db example.db -o dump.sql

positional arguments:
  COMMAND
    action to
    perform
  TARGET
    zero or
    more
    arguments,
    meaning
    depends on
    command

optional arguments:
  -h
    show this
    help
    message and
    exit
  --help
    show this
    help
    message and
    exit

main options:
  -f FILE, --file FILE
    load a list
    of targets
    from a
    plain text
    file
  --config FILE
    global conf
    iguration
    file
  --user-config FILE
    per-user co
    nfiguration
    file
  -p NAME, --profile NAME
    profile to
    use
  --ui-mode MODE
    UI mode
  -v, --verbose
    increase
    output
    verbosity
  -q, --quiet
    suppress
    text output
  --color
    use colors
    in console
    output
  --no-color
    suppress
    colors in
    console
    output

audit options:
  --audit-name NAME
    customize
    the audit
    name
  -db DATABASE, --audit-db DATABASE
    specify a
    database
    filename
  -nd, --no-db
    do not
    store the
    results in
    a database
  -i FILENAME, --input FILENAME
    read
    results
    from
    external
    tools right
    before the
    audit
  -ni, --no-input
    do not read
    results
    from
    external
    tools

report options:
  -o FILENAME, --output FILENAME
    write the
    results of
    the audit
    to this
    file (use -
    for stdout)
  -no, --no-output
    do not
    output the
    results
  --full
    produce
    fully
    detailed
    reports
  --brief
    report only
    the
    highlights

network options:
  --allow-subdomains
    include
    subdomains
    in the
    target
    scope
  --forbid-subdomains
    do not
    include
    subdomains
    in the
    target
    scope
  --parent
    include
    parent
    folders in
    the target
    scope
  -np, --no-parent
    do not
    include
    parent
    folders in
    the target
    scope
  -r DEPTH, --depth DEPTH
    maximum
    spidering
    depth (use
    "infinite"
    for no
    limit)
  --follow-redirects
    follow
    redirects
  --no-follow-redirects
    do not
    follow
    redirects
  --follow-first
    always
    follow a
    redirection
    on the
    target URL
    itself
  --no-follow-first
    don't treat
    a
    redirection
    on a target
    URL as a
    special
    case
  --max-connections MAX_CONNECTIONS
    maximum
    number of
    concurrent
    connections
    per host
  -l MAX_LINKS, --max-links MAX_LINKS
    maximum
    number of
    links to
    analyze (0
    =>
    infinite)
  -pu USER, --proxy-user USER
    HTTP proxy
    username
  -pp PASS, --proxy-pass PASS
    HTTP proxy
    password
  -pa ADDRESS, --proxy-addr ADDRESS
    HTTP proxy
    address
  -pn PORT, --proxy-port PORT
    HTTP proxy
    port number
  --cookie COOKIE
    set cookie
    for
    requests
  --user-agent USER_AGENT
    set a
    custom user
    agent or
    'random'
    value
  --cookie-file FILE
    load a
    cookie from
    file
  --persistent-cache
    use a
    persistent
    network
    cache
    [default]
  --volatile-cache
    use a
    volatile
    network
    cache

plugin options:
  -a PLUGIN:KEY=VALUE, --plugin-arg PLUGIN:KEY=VALUE
    pass an
    argument to
    a plugin
  -e PLUGIN, --enable-plugin PLUGIN
    enable a

```