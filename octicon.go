package octicon

import (
	"fmt"
	"html/template"
	"strings"
)

// Icon returns a string representing the named Octicon.
func Icon(name string, height int, extraClasses ...string) (string, bool) {
	switch name {
	case "accessibility":
		return Accessibility(height, extraClasses...)
	case "alert":
		return Alert(height, extraClasses...)
	case "alert-fill":
		return AlertFill(height, extraClasses...)
	case "apps":
		return Apps(height, extraClasses...)
	case "archive":
		return Archive(height, extraClasses...)
	case "arrow-both":
		return ArrowBoth(height, extraClasses...)
	case "arrow-down":
		return ArrowDown(height, extraClasses...)
	case "arrow-down-left":
		return ArrowDownLeft(height, extraClasses...)
	case "arrow-down-right":
		return ArrowDownRight(height, extraClasses...)
	case "arrow-left":
		return ArrowLeft(height, extraClasses...)
	case "arrow-right":
		return ArrowRight(height, extraClasses...)
	case "arrow-switch":
		return ArrowSwitch(height, extraClasses...)
	case "arrow-up":
		return ArrowUp(height, extraClasses...)
	case "arrow-up-left":
		return ArrowUpLeft(height, extraClasses...)
	case "arrow-up-right":
		return ArrowUpRight(height, extraClasses...)
	case "beaker":
		return Beaker(height, extraClasses...)
	case "bell":
		return Bell(height, extraClasses...)
	case "bell-fill":
		return BellFill(height, extraClasses...)
	case "bell-slash":
		return BellSlash(height, extraClasses...)
	case "blocked":
		return Blocked(height, extraClasses...)
	case "bold":
		return Bold(height, extraClasses...)
	case "book":
		return Book(height, extraClasses...)
	case "bookmark":
		return Bookmark(height, extraClasses...)
	case "bookmark-fill":
		return BookmarkFill(height, extraClasses...)
	case "bookmark-slash":
		return BookmarkSlash(height, extraClasses...)
	case "bookmark-slash-fill":
		return BookmarkSlashFill(height, extraClasses...)
	case "briefcase":
		return Briefcase(height, extraClasses...)
	case "broadcast":
		return Broadcast(height, extraClasses...)
	case "browser":
		return Browser(height, extraClasses...)
	case "bug":
		return Bug(height, extraClasses...)
	case "calendar":
		return Calendar(height, extraClasses...)
	case "check":
		return Check(height, extraClasses...)
	case "check-circle":
		return CheckCircle(height, extraClasses...)
	case "check-circle-fill":
		return CheckCircleFill(height, extraClasses...)
	case "checklist":
		return Checklist(height, extraClasses...)
	case "chevron-down":
		return ChevronDown(height, extraClasses...)
	case "chevron-left":
		return ChevronLeft(height, extraClasses...)
	case "chevron-right":
		return ChevronRight(height, extraClasses...)
	case "chevron-up":
		return ChevronUp(height, extraClasses...)
	case "circle":
		return Circle(height, extraClasses...)
	case "circle-slash":
		return CircleSlash(height, extraClasses...)
	case "clock":
		return Clock(height, extraClasses...)
	case "cloud":
		return Cloud(height, extraClasses...)
	case "cloud-offline":
		return CloudOffline(height, extraClasses...)
	case "code":
		return Code(height, extraClasses...)
	case "code-of-conduct":
		return CodeOfConduct(height, extraClasses...)
	case "code-review":
		return CodeReview(height, extraClasses...)
	case "code-square":
		return CodeSquare(height, extraClasses...)
	case "codescan":
		return Codescan(height, extraClasses...)
	case "codescan-checkmark":
		return CodescanCheckmark(height, extraClasses...)
	case "codespaces":
		return Codespaces(height, extraClasses...)
	case "columns":
		return Columns(height, extraClasses...)
	case "comment":
		return Comment(height, extraClasses...)
	case "comment-discussion":
		return CommentDiscussion(height, extraClasses...)
	case "commit":
		return Commit(height, extraClasses...)
	case "container":
		return Container(height, extraClasses...)
	case "copy":
		return Copy(height, extraClasses...)
	case "cpu":
		return Cpu(height, extraClasses...)
	case "credit-card":
		return CreditCard(height, extraClasses...)
	case "cross-reference":
		return CrossReference(height, extraClasses...)
	case "dash":
		return Dash(height, extraClasses...)
	case "database":
		return Database(height, extraClasses...)
	case "dependabot":
		return Dependabot(height, extraClasses...)
	case "desktop-download":
		return DesktopDownload(height, extraClasses...)
	case "device-camera":
		return DeviceCamera(height, extraClasses...)
	case "device-camera-video":
		return DeviceCameraVideo(height, extraClasses...)
	case "device-desktop":
		return DeviceDesktop(height, extraClasses...)
	case "device-mobile":
		return DeviceMobile(height, extraClasses...)
	case "diamond":
		return Diamond(height, extraClasses...)
	case "diff":
		return Diff(height, extraClasses...)
	case "diff-added":
		return DiffAdded(height, extraClasses...)
	case "diff-ignored":
		return DiffIgnored(height, extraClasses...)
	case "diff-modified":
		return DiffModified(height, extraClasses...)
	case "diff-removed":
		return DiffRemoved(height, extraClasses...)
	case "diff-renamed":
		return DiffRenamed(height, extraClasses...)
	case "dot":
		return Dot(height, extraClasses...)
	case "dot-fill":
		return DotFill(height, extraClasses...)
	case "download":
		return Download(height, extraClasses...)
	case "duplicate":
		return Duplicate(height, extraClasses...)
	case "ellipsis":
		return Ellipsis(height, extraClasses...)
	case "eye":
		return Eye(height, extraClasses...)
	case "eye-closed":
		return EyeClosed(height, extraClasses...)
	case "feed-discussion":
		return FeedDiscussion(height, extraClasses...)
	case "feed-forked":
		return FeedForked(height, extraClasses...)
	case "feed-heart":
		return FeedHeart(height, extraClasses...)
	case "feed-merged":
		return FeedMerged(height, extraClasses...)
	case "feed-person":
		return FeedPerson(height, extraClasses...)
	case "feed-repo":
		return FeedRepo(height, extraClasses...)
	case "feed-rocket":
		return FeedRocket(height, extraClasses...)
	case "feed-star":
		return FeedStar(height, extraClasses...)
	case "feed-tag":
		return FeedTag(height, extraClasses...)
	case "feed-trophy":
		return FeedTrophy(height, extraClasses...)
	case "file":
		return File(height, extraClasses...)
	case "file-added":
		return FileAdded(height, extraClasses...)
	case "file-badge":
		return FileBadge(height, extraClasses...)
	case "file-binary":
		return FileBinary(height, extraClasses...)
	case "file-code":
		return FileCode(height, extraClasses...)
	case "file-diff":
		return FileDiff(height, extraClasses...)
	case "file-directory":
		return FileDirectory(height, extraClasses...)
	case "file-directory-fill":
		return FileDirectoryFill(height, extraClasses...)
	case "file-directory-open-fill":
		return FileDirectoryOpenFill(height, extraClasses...)
	case "file-media":
		return FileMedia(height, extraClasses...)
	case "file-moved":
		return FileMoved(height, extraClasses...)
	case "file-removed":
		return FileRemoved(height, extraClasses...)
	case "file-submodule":
		return FileSubmodule(height, extraClasses...)
	case "file-symlink-file":
		return FileSymlinkFile(height, extraClasses...)
	case "file-zip":
		return FileZip(height, extraClasses...)
	case "filter":
		return Filter(height, extraClasses...)
	case "flame":
		return Flame(height, extraClasses...)
	case "fold":
		return Fold(height, extraClasses...)
	case "fold-down":
		return FoldDown(height, extraClasses...)
	case "fold-up":
		return FoldUp(height, extraClasses...)
	case "gear":
		return Gear(height, extraClasses...)
	case "gift":
		return Gift(height, extraClasses...)
	case "git-branch":
		return GitBranch(height, extraClasses...)
	case "git-commit":
		return GitCommit(height, extraClasses...)
	case "git-compare":
		return GitCompare(height, extraClasses...)
	case "git-merge":
		return GitMerge(height, extraClasses...)
	case "git-pull-request":
		return GitPullRequest(height, extraClasses...)
	case "git-pull-request-closed":
		return GitPullRequestClosed(height, extraClasses...)
	case "git-pull-request-draft":
		return GitPullRequestDraft(height, extraClasses...)
	case "globe":
		return Globe(height, extraClasses...)
	case "grabber":
		return Grabber(height, extraClasses...)
	case "graph":
		return Graph(height, extraClasses...)
	case "hash":
		return Hash(height, extraClasses...)
	case "heading":
		return Heading(height, extraClasses...)
	case "heart":
		return Heart(height, extraClasses...)
	case "heart-fill":
		return HeartFill(height, extraClasses...)
	case "history":
		return History(height, extraClasses...)
	case "home":
		return Home(height, extraClasses...)
	case "home-fill":
		return HomeFill(height, extraClasses...)
	case "horizontal-rule":
		return HorizontalRule(height, extraClasses...)
	case "hourglass":
		return Hourglass(height, extraClasses...)
	case "hubot":
		return Hubot(height, extraClasses...)
	case "id-badge":
		return IdBadge(height, extraClasses...)
	case "image":
		return Image(height, extraClasses...)
	case "inbox":
		return Inbox(height, extraClasses...)
	case "infinity":
		return Infinity(height, extraClasses...)
	case "info":
		return Info(height, extraClasses...)
	case "issue-closed":
		return IssueClosed(height, extraClasses...)
	case "issue-draft":
		return IssueDraft(height, extraClasses...)
	case "issue-opened":
		return IssueOpened(height, extraClasses...)
	case "issue-reopened":
		return IssueReopened(height, extraClasses...)
	case "italic":
		return Italic(height, extraClasses...)
	case "iterations":
		return Iterations(height, extraClasses...)
	case "kebab-horizontal":
		return KebabHorizontal(height, extraClasses...)
	case "key":
		return Key(height, extraClasses...)
	case "key-asterisk":
		return KeyAsterisk(height, extraClasses...)
	case "law":
		return Law(height, extraClasses...)
	case "light-bulb":
		return LightBulb(height, extraClasses...)
	case "link":
		return Link(height, extraClasses...)
	case "link-external":
		return LinkExternal(height, extraClasses...)
	case "list-ordered":
		return ListOrdered(height, extraClasses...)
	case "list-unordered":
		return ListUnordered(height, extraClasses...)
	case "location":
		return Location(height, extraClasses...)
	case "lock":
		return Lock(height, extraClasses...)
	case "log":
		return Log(height, extraClasses...)
	case "logo-gist":
		return LogoGist(height, extraClasses...)
	case "logo-github":
		return LogoGithub(height, extraClasses...)
	case "mail":
		return Mail(height, extraClasses...)
	case "mark-github":
		return MarkGithub(height, extraClasses...)
	case "markdown":
		return Markdown(height, extraClasses...)
	case "megaphone":
		return Megaphone(height, extraClasses...)
	case "mention":
		return Mention(height, extraClasses...)
	case "meter":
		return Meter(height, extraClasses...)
	case "milestone":
		return Milestone(height, extraClasses...)
	case "mirror":
		return Mirror(height, extraClasses...)
	case "moon":
		return Moon(height, extraClasses...)
	case "mortar-board":
		return MortarBoard(height, extraClasses...)
	case "multi-select":
		return MultiSelect(height, extraClasses...)
	case "mute":
		return Mute(height, extraClasses...)
	case "no-entry":
		return NoEntry(height, extraClasses...)
	case "no-entry-fill":
		return NoEntryFill(height, extraClasses...)
	case "north-star":
		return NorthStar(height, extraClasses...)
	case "note":
		return Note(height, extraClasses...)
	case "number":
		return Number(height, extraClasses...)
	case "organization":
		return Organization(height, extraClasses...)
	case "package":
		return Package(height, extraClasses...)
	case "package-dependencies":
		return PackageDependencies(height, extraClasses...)
	case "package-dependents":
		return PackageDependents(height, extraClasses...)
	case "paintbrush":
		return Paintbrush(height, extraClasses...)
	case "paper-airplane":
		return PaperAirplane(height, extraClasses...)
	case "paste":
		return Paste(height, extraClasses...)
	case "pencil":
		return Pencil(height, extraClasses...)
	case "people":
		return People(height, extraClasses...)
	case "person":
		return Person(height, extraClasses...)
	case "person-add":
		return PersonAdd(height, extraClasses...)
	case "person-fill":
		return PersonFill(height, extraClasses...)
	case "pin":
		return Pin(height, extraClasses...)
	case "play":
		return Play(height, extraClasses...)
	case "plug":
		return Plug(height, extraClasses...)
	case "plus":
		return Plus(height, extraClasses...)
	case "plus-circle":
		return PlusCircle(height, extraClasses...)
	case "project":
		return Project(height, extraClasses...)
	case "pulse":
		return Pulse(height, extraClasses...)
	case "question":
		return Question(height, extraClasses...)
	case "quote":
		return Quote(height, extraClasses...)
	case "reply":
		return Reply(height, extraClasses...)
	case "repo":
		return Repo(height, extraClasses...)
	case "repo-clone":
		return RepoClone(height, extraClasses...)
	case "repo-deleted":
		return RepoDeleted(height, extraClasses...)
	case "repo-forked":
		return RepoForked(height, extraClasses...)
	case "repo-locked":
		return RepoLocked(height, extraClasses...)
	case "repo-pull":
		return RepoPull(height, extraClasses...)
	case "repo-push":
		return RepoPush(height, extraClasses...)
	case "repo-template":
		return RepoTemplate(height, extraClasses...)
	case "report":
		return Report(height, extraClasses...)
	case "rocket":
		return Rocket(height, extraClasses...)
	case "rows":
		return Rows(height, extraClasses...)
	case "rss":
		return Rss(height, extraClasses...)
	case "ruby":
		return Ruby(height, extraClasses...)
	case "screen-full":
		return ScreenFull(height, extraClasses...)
	case "screen-normal":
		return ScreenNormal(height, extraClasses...)
	case "search":
		return Search(height, extraClasses...)
	case "server":
		return Server(height, extraClasses...)
	case "share":
		return Share(height, extraClasses...)
	case "share-android":
		return ShareAndroid(height, extraClasses...)
	case "shield":
		return Shield(height, extraClasses...)
	case "shield-check":
		return ShieldCheck(height, extraClasses...)
	case "shield-lock":
		return ShieldLock(height, extraClasses...)
	case "shield-x":
		return ShieldX(height, extraClasses...)
	case "sidebar-collapse":
		return SidebarCollapse(height, extraClasses...)
	case "sidebar-expand":
		return SidebarExpand(height, extraClasses...)
	case "sign-in":
		return SignIn(height, extraClasses...)
	case "sign-out":
		return SignOut(height, extraClasses...)
	case "single-select":
		return SingleSelect(height, extraClasses...)
	case "skip":
		return Skip(height, extraClasses...)
	case "sliders":
		return Sliders(height, extraClasses...)
	case "smiley":
		return Smiley(height, extraClasses...)
	case "sort-asc":
		return SortAsc(height, extraClasses...)
	case "sort-desc":
		return SortDesc(height, extraClasses...)
	case "square":
		return Square(height, extraClasses...)
	case "square-fill":
		return SquareFill(height, extraClasses...)
	case "squirrel":
		return Squirrel(height, extraClasses...)
	case "stack":
		return Stack(height, extraClasses...)
	case "star":
		return Star(height, extraClasses...)
	case "star-fill":
		return StarFill(height, extraClasses...)
	case "stop":
		return Stop(height, extraClasses...)
	case "stopwatch":
		return Stopwatch(height, extraClasses...)
	case "strikethrough":
		return Strikethrough(height, extraClasses...)
	case "sun":
		return Sun(height, extraClasses...)
	case "sync":
		return Sync(height, extraClasses...)
	case "tab":
		return Tab(height, extraClasses...)
	case "tab-external":
		return TabExternal(height, extraClasses...)
	case "table":
		return Table(height, extraClasses...)
	case "tag":
		return Tag(height, extraClasses...)
	case "tasklist":
		return Tasklist(height, extraClasses...)
	case "telescope":
		return Telescope(height, extraClasses...)
	case "telescope-fill":
		return TelescopeFill(height, extraClasses...)
	case "terminal":
		return Terminal(height, extraClasses...)
	case "three-bars":
		return ThreeBars(height, extraClasses...)
	case "thumbsdown":
		return Thumbsdown(height, extraClasses...)
	case "thumbsup":
		return Thumbsup(height, extraClasses...)
	case "tools":
		return Tools(height, extraClasses...)
	case "trash":
		return Trash(height, extraClasses...)
	case "triangle-down":
		return TriangleDown(height, extraClasses...)
	case "triangle-left":
		return TriangleLeft(height, extraClasses...)
	case "triangle-right":
		return TriangleRight(height, extraClasses...)
	case "triangle-up":
		return TriangleUp(height, extraClasses...)
	case "trophy":
		return Trophy(height, extraClasses...)
	case "typography":
		return Typography(height, extraClasses...)
	case "unfold":
		return Unfold(height, extraClasses...)
	case "unlock":
		return Unlock(height, extraClasses...)
	case "unmute":
		return Unmute(height, extraClasses...)
	case "unverified":
		return Unverified(height, extraClasses...)
	case "upload":
		return Upload(height, extraClasses...)
	case "verified":
		return Verified(height, extraClasses...)
	case "versions":
		return Versions(height, extraClasses...)
	case "video":
		return Video(height, extraClasses...)
	case "webhook":
		return Webhook(height, extraClasses...)
	case "workflow":
		return Workflow(height, extraClasses...)
	case "x":
		return X(height, extraClasses...)
	case "x-circle":
		return XCircle(height, extraClasses...)
	case "x-circle-fill":
		return XCircleFill(height, extraClasses...)
	case "zap":
		return Zap(height, extraClasses...)
	default:
		return "", false
	}
}

func IconTemplateFunc(name string, height int, extraClasses ...string) (template.HTML, error) {
	i, ok := Icon(name, height, extraClasses...)
	if !ok {
		return "", fmt.Errorf("unknown icon (%s) or height (%d)", name, height)
	}

	return template.HTML(i), nil
}

// Accessibility returns a string representing an "accessibility" Octicon.
func Accessibility(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-accessibility%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.923 5.302a3 3 0 10-3.847 0A2.713 2.713 0 005.9 5.5H2A.75.75 0 002 7h3.3l-.578 5.163-.362 2.997a.75.75 0 101.49.18L6.132 13h3.736l.282 2.34a.75.75 0 101.49-.18l-.362-2.997L10.7 7H14a.75.75 0 000-1.5h-3.899a2.697 2.697 0 00-.178-.198zM9.5 3a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zm-.3 4.073l.495 4.427h-3.39l.496-4.427a1.207 1.207 0 012.398 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Alert returns a string representing an "alert" Octicon.
func Alert(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-alert%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.22 1.754a.25.25 0 00-.44 0L1.698 13.132a.25.25 0 00.22.368h12.164a.25.25 0 00.22-.368L8.22 1.754zm-1.763-.707c.659-1.234 2.427-1.234 3.086 0l6.082 11.378A1.75 1.75 0 0114.082 15H1.918a1.75 1.75 0 01-1.543-2.575L6.457 1.047zM9 11a1 1 0 11-2 0 1 1 0 012 0zm-.25-5.25a.75.75 0 00-1.5 0v2.5a.75.75 0 001.5 0v-2.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-alert%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M13 17.5a1 1 0 11-2 0 1 1 0 012 0zm-.25-8.25a.75.75 0 00-1.5 0v4.5a.75.75 0 001.5 0v-4.5z\"></path><path fill-rule=\"evenodd\" d=\"M9.836 3.244c.963-1.665 3.365-1.665 4.328 0l8.967 15.504c.963 1.667-.24 3.752-2.165 3.752H3.034c-1.926 0-3.128-2.085-2.165-3.752L9.836 3.244zm3.03.751a1 1 0 00-1.732 0L2.168 19.499A1 1 0 003.034 21h17.932a1 1 0 00.866-1.5L12.866 3.994z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// AlertFill returns a string representing an "alert-fill" Octicon.
func AlertFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 12:
		return fmt.Sprintf("<svg class=\"octicon octicon-alert-fill%s\" height=\"12\" width=\"12\" viewbox=\"0 0 12 12\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.855.708c.5-.896 1.79-.896 2.29 0l4.675 8.351a1.312 1.312 0 01-1.146 1.954H1.33A1.312 1.312 0 01.183 9.058L4.855.708zM7 7V3H5v4h2zm-1 3a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Apps returns a string representing an "apps" Octicon.
func Apps(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-apps%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 3.25c0-.966.784-1.75 1.75-1.75h2.5c.966 0 1.75.784 1.75 1.75v2.5A1.75 1.75 0 015.75 7.5h-2.5A1.75 1.75 0 011.5 5.75v-2.5zM3.25 3a.25.25 0 00-.25.25v2.5c0 .138.112.25.25.25h2.5A.25.25 0 006 5.75v-2.5A.25.25 0 005.75 3h-2.5zM1.5 10.25c0-.966.784-1.75 1.75-1.75h2.5c.966 0 1.75.784 1.75 1.75v2.5a1.75 1.75 0 01-1.75 1.75h-2.5a1.75 1.75 0 01-1.75-1.75v-2.5zM3.25 10a.25.25 0 00-.25.25v2.5c0 .138.112.25.25.25h2.5a.25.25 0 00.25-.25v-2.5a.25.25 0 00-.25-.25h-2.5zM8.5 3.25c0-.966.784-1.75 1.75-1.75h2.5c.966 0 1.75.784 1.75 1.75v2.5a1.75 1.75 0 01-1.75 1.75h-2.5A1.75 1.75 0 018.5 5.75v-2.5zM10.25 3a.25.25 0 00-.25.25v2.5c0 .138.112.25.25.25h2.5a.25.25 0 00.25-.25v-2.5a.25.25 0 00-.25-.25h-2.5zM8.5 10.25c0-.966.784-1.75 1.75-1.75h2.5c.966 0 1.75.784 1.75 1.75v2.5a1.75 1.75 0 01-1.75 1.75h-2.5a1.75 1.75 0 01-1.75-1.75v-2.5zm1.75-.25a.25.25 0 00-.25.25v2.5c0 .138.112.25.25.25h2.5a.25.25 0 00.25-.25v-2.5a.25.25 0 00-.25-.25h-2.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Archive returns a string representing an "archive" Octicon.
func Archive(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-archive%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 2.5a.25.25 0 00-.25.25v1.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25v-1.5a.25.25 0 00-.25-.25H1.75zM0 2.75C0 1.784.784 1 1.75 1h12.5c.966 0 1.75.784 1.75 1.75v1.5A1.75 1.75 0 0114.25 6H1.75A1.75 1.75 0 010 4.25v-1.5zM1.75 7a.75.75 0 01.75.75v5.5c0 .138.112.25.25.25h10.5a.25.25 0 00.25-.25v-5.5a.75.75 0 111.5 0v5.5A1.75 1.75 0 0113.25 15H2.75A1.75 1.75 0 011 13.25v-5.5A.75.75 0 011.75 7zm4.5 1a.75.75 0 000 1.5h3.5a.75.75 0 100-1.5h-3.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-archive%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 2A1.75 1.75 0 001 3.75v3.5C1 8.216 1.784 9 2.75 9h18.5A1.75 1.75 0 0023 7.25v-3.5A1.75 1.75 0 0021.25 2H2.75zm18.5 1.5H2.75a.25.25 0 00-.25.25v3.5c0 .138.112.25.25.25h18.5a.25.25 0 00.25-.25v-3.5a.25.25 0 00-.25-.25z\"></path><path d=\"M2.75 10a.75.75 0 01.75.75v9.5c0 .138.112.25.25.25h16.5a.25.25 0 00.25-.25v-9.5a.75.75 0 011.5 0v9.5A1.75 1.75 0 0120.25 22H3.75A1.75 1.75 0 012 20.25v-9.5a.75.75 0 01.75-.75z\"></path><path d=\"M9.75 11.5a.75.75 0 000 1.5h4.5a.75.75 0 000-1.5h-4.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowBoth returns a string representing an "arrow-both" Octicon.
func ArrowBoth(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-both%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.72 3.72a.75.75 0 011.06 1.06L2.56 7h10.88l-2.22-2.22a.75.75 0 011.06-1.06l3.5 3.5a.75.75 0 010 1.06l-3.5 3.5a.75.75 0 11-1.06-1.06l2.22-2.22H2.56l2.22 2.22a.75.75 0 11-1.06 1.06l-3.5-3.5a.75.75 0 010-1.06l3.5-3.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-both%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M7.78 5.97a.75.75 0 00-1.06 0l-5.25 5.25a.75.75 0 000 1.06l5.25 5.25a.75.75 0 001.06-1.06L3.81 12.5h16.38l-3.97 3.97a.75.75 0 101.06 1.06l5.25-5.25a.75.75 0 000-1.06l-5.25-5.25a.75.75 0 10-1.06 1.06L20.19 11H3.81l3.97-3.97a.75.75 0 000-1.06z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowDown returns a string representing an "arrow-down" Octicon.
func ArrowDown(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-down%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13.03 8.22a.75.75 0 010 1.06l-4.25 4.25a.75.75 0 01-1.06 0L3.47 9.28a.75.75 0 011.06-1.06l2.97 2.97V3.75a.75.75 0 011.5 0v7.44l2.97-2.97a.75.75 0 011.06 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-down%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.97 13.22a.75.75 0 000 1.06l6.25 6.25a.75.75 0 001.06 0l6.25-6.25a.75.75 0 10-1.06-1.06l-4.97 4.97V3.75a.75.75 0 00-1.5 0v14.44l-4.97-4.97a.75.75 0 00-1.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowDownLeft returns a string representing an "arrow-down-left" Octicon.
func ArrowDownLeft(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-down-left%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.75 8.5a.75.75 0 00-.75.75v9c0 .414.336.75.75.75h9a.75.75 0 000-1.5H7.56L17.78 7.28a.75.75 0 00-1.06-1.06L6.5 16.44V9.25a.75.75 0 00-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowDownRight returns a string representing an "arrow-down-right" Octicon.
func ArrowDownRight(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-down-right%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M18.25 8.5a.75.75 0 01.75.75v9a.75.75 0 01-.75.75h-9a.75.75 0 010-1.5h7.19L6.22 7.28a.75.75 0 011.06-1.06L17.5 16.44V9.25a.75.75 0 01.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowLeft returns a string representing an "arrow-left" Octicon.
func ArrowLeft(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-left%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.78 12.53a.75.75 0 01-1.06 0L2.47 8.28a.75.75 0 010-1.06l4.25-4.25a.75.75 0 011.06 1.06L4.81 7h7.44a.75.75 0 010 1.5H4.81l2.97 2.97a.75.75 0 010 1.06z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-left%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.78 19.03a.75.75 0 01-1.06 0l-6.25-6.25a.75.75 0 010-1.06l6.25-6.25a.75.75 0 111.06 1.06L5.81 11.5h14.44a.75.75 0 010 1.5H5.81l4.97 4.97a.75.75 0 010 1.06z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowRight returns a string representing an "arrow-right" Octicon.
func ArrowRight(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-right%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.22 2.97a.75.75 0 011.06 0l4.25 4.25a.75.75 0 010 1.06l-4.25 4.25a.75.75 0 01-1.06-1.06l2.97-2.97H3.75a.75.75 0 010-1.5h7.44L8.22 4.03a.75.75 0 010-1.06z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-right%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13.22 19.03a.75.75 0 001.06 0l6.25-6.25a.75.75 0 000-1.06l-6.25-6.25a.75.75 0 10-1.06 1.06l4.97 4.97H3.75a.75.75 0 000 1.5h14.44l-4.97 4.97a.75.75 0 000 1.06z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowSwitch returns a string representing an "arrow-switch" Octicon.
func ArrowSwitch(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-switch%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M5.22 14.78a.75.75 0 001.06-1.06L4.56 12h8.69a.75.75 0 000-1.5H4.56l1.72-1.72a.75.75 0 00-1.06-1.06l-3 3a.75.75 0 000 1.06l3 3zm5.56-6.5a.75.75 0 11-1.06-1.06l1.72-1.72H2.75a.75.75 0 010-1.5h8.69L9.72 2.28a.75.75 0 011.06-1.06l3 3a.75.75 0 010 1.06l-3 3z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-switch%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M7.72 21.78a.75.75 0 001.06-1.06L5.56 17.5h14.69a.75.75 0 000-1.5H5.56l3.22-3.22a.75.75 0 10-1.06-1.06l-4.5 4.5a.75.75 0 000 1.06l4.5 4.5zm8.56-9.5a.75.75 0 11-1.06-1.06L18.44 8H3.75a.75.75 0 010-1.5h14.69l-3.22-3.22a.75.75 0 011.06-1.06l4.5 4.5a.75.75 0 010 1.06l-4.5 4.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowUp returns a string representing an "arrow-up" Octicon.
func ArrowUp(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-up%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.47 7.78a.75.75 0 010-1.06l4.25-4.25a.75.75 0 011.06 0l4.25 4.25a.75.75 0 01-1.06 1.06L9 4.81v7.44a.75.75 0 01-1.5 0V4.81L4.53 7.78a.75.75 0 01-1.06 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-up%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M18.655 10.405a.75.75 0 000-1.06l-6.25-6.25a.75.75 0 00-1.06 0l-6.25 6.25a.75.75 0 101.06 1.06l4.97-4.97v14.44a.75.75 0 001.5 0V5.435l4.97 4.97a.75.75 0 001.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowUpLeft returns a string representing an "arrow-up-left" Octicon.
func ArrowUpLeft(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-up-left%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.75 15.5a.75.75 0 01-.75-.75v-9A.75.75 0 015.75 5h9a.75.75 0 010 1.5H7.56l10.22 10.22a.75.75 0 11-1.06 1.06L6.5 7.56v7.19a.75.75 0 01-.75.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ArrowUpRight returns a string representing an "arrow-up-right" Octicon.
func ArrowUpRight(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-arrow-up-right%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M18.25 15.5a.75.75 0 00.75-.75v-9a.75.75 0 00-.75-.75h-9a.75.75 0 000 1.5h7.19L6.22 16.72a.75.75 0 101.06 1.06L17.5 7.56v7.19c0 .414.336.75.75.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Beaker returns a string representing an "beaker" Octicon.
func Beaker(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-beaker%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5 5.782V2.5h-.25a.75.75 0 010-1.5h6.5a.75.75 0 010 1.5H11v3.282l3.666 5.76C15.619 13.04 14.543 15 12.767 15H3.233c-1.776 0-2.852-1.96-1.899-3.458L5 5.782zM9.5 2.5h-3V6a.75.75 0 01-.117.403L4.73 9h6.54L9.617 6.403A.75.75 0 019.5 6V2.5zm-6.9 9.847L3.775 10.5h8.45l1.175 1.847a.75.75 0 01-.633 1.153H3.233a.75.75 0 01-.633-1.153z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-beaker%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 8.807V3.5h-.563a.75.75 0 010-1.5h9.125a.75.75 0 010 1.5H16v5.307l5.125 9.301c.964 1.75-.302 3.892-2.299 3.892H5.174c-1.998 0-3.263-2.142-2.3-3.892L8 8.807zM14.5 3.5h-5V9a.75.75 0 01-.093.362L7.127 13.5h9.746l-2.28-4.138A.75.75 0 0114.5 9V3.5zM4.189 18.832L6.3 15h11.4l2.111 3.832a1.125 1.125 0 01-.985 1.668H5.174a1.125 1.125 0 01-.985-1.668z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Bell returns a string representing an "bell" Octicon.
func Bell(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-bell%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8 16a2 2 0 001.985-1.75c.017-.137-.097-.25-.235-.25h-3.5c-.138 0-.252.113-.235.25A2 2 0 008 16z\"></path><path fill-rule=\"evenodd\" d=\"M8 1.5A3.5 3.5 0 004.5 5v2.947c0 .346-.102.683-.294.97l-1.703 2.556a.018.018 0 00-.003.01l.001.006c0 .002.002.004.004.006a.017.017 0 00.006.004l.007.001h10.964l.007-.001a.016.016 0 00.006-.004.016.016 0 00.004-.006l.001-.007a.017.017 0 00-.003-.01l-1.703-2.554a1.75 1.75 0 01-.294-.97V5A3.5 3.5 0 008 1.5zM3 5a5 5 0 0110 0v2.947c0 .05.015.098.042.139l1.703 2.555A1.518 1.518 0 0113.482 13H2.518a1.518 1.518 0 01-1.263-2.36l1.703-2.554A.25.25 0 003 7.947V5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bell%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 1C8.318 1 5 3.565 5 7v4.539a3.25 3.25 0 01-.546 1.803l-2.2 3.299A1.518 1.518 0 003.519 19H8.5a3.5 3.5 0 107 0h4.982a1.518 1.518 0 001.263-2.36l-2.2-3.298A3.25 3.25 0 0119 11.539V7c0-3.435-3.319-6-7-6zM6.5 7c0-2.364 2.383-4.5 5.5-4.5s5.5 2.136 5.5 4.5v4.539c0 .938.278 1.854.798 2.635l2.199 3.299a.017.017 0 01.003.01l-.001.006-.004.006-.006.004-.007.001H3.518l-.007-.001-.006-.004-.004-.006-.001-.007.003-.01 2.2-3.298a4.75 4.75 0 00.797-2.635V7zM14 19h-4a2 2 0 104 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// BellFill returns a string representing an "bell-fill" Octicon.
func BellFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-bell-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8 16c.9 0 1.7-.6 1.9-1.5.1-.3-.1-.5-.4-.5h-3c-.3 0-.5.2-.4.5.2.9 1 1.5 1.9 1.5zM3 5c0-2.8 2.2-5 5-5s5 2.2 5 5v3l1.7 2.6c.2.2.3.5.3.8 0 .8-.7 1.5-1.5 1.5h-11c-.8.1-1.5-.6-1.5-1.4 0-.3.1-.6.3-.8L3 8.1V5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bell-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 8a6 6 0 1112 0v2.917c0 .703.228 1.387.65 1.95L20.7 15.6a1.5 1.5 0 01-1.2 2.4h-15a1.5 1.5 0 01-1.2-2.4l2.05-2.733a3.25 3.25 0 00.65-1.95V8zm6 13.5A3.502 3.502 0 018.645 19h6.71A3.502 3.502 0 0112 21.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// BellSlash returns a string representing an "bell-slash" Octicon.
func BellSlash(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-bell-slash%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 1.5c-.997 0-1.895.416-2.534 1.086A.75.75 0 014.38 1.55 5 5 0 0113 5v2.373a.75.75 0 01-1.5 0V5A3.5 3.5 0 008 1.5zM4.182 4.31L1.19 2.143a.75.75 0 10-.88 1.214L3 5.305v2.642a.25.25 0 01-.042.139L1.255 10.64A1.518 1.518 0 002.518 13h11.108l1.184.857a.75.75 0 10.88-1.214l-1.375-.996a1.196 1.196 0 00-.013-.01L4.198 4.321a.733.733 0 00-.016-.011zm7.373 7.19L4.5 6.391v1.556c0 .346-.102.683-.294.97l-1.703 2.556a.018.018 0 00-.003.01.015.015 0 00.005.012.017.017 0 00.006.004l.007.001h9.037zM8 16a2 2 0 001.985-1.75c.017-.137-.097-.25-.235-.25h-3.5c-.138 0-.252.113-.235.25A2 2 0 008 16z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bell-slash%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.22 1.22a.75.75 0 011.06 0l20.5 20.5a.75.75 0 11-1.06 1.06L17.94 19H15.5a3.5 3.5 0 11-7 0H3.518a1.518 1.518 0 01-1.263-2.36l2.2-3.298A3.25 3.25 0 005 11.539V7c0-.294.025-.583.073-.866L1.22 2.28a.75.75 0 010-1.06zM10 19a2 2 0 104 0h-4zM6.5 7.56l9.94 9.94H3.517l-.007-.001-.006-.004-.004-.006-.001-.007.003-.01 2.2-3.298a4.75 4.75 0 00.797-2.635V7.56z\"></path><path d=\"M12 2.5c-1.463 0-2.8.485-3.788 1.257l-.04.032a.75.75 0 11-.935-1.173l.05-.04C8.548 1.59 10.212 1 12 1c3.681 0 7 2.565 7 6v4.539c0 .642.19 1.269.546 1.803l1.328 1.992a.75.75 0 11-1.248.832l-1.328-1.992a4.75 4.75 0 01-.798-2.635V7c0-2.364-2.383-4.5-5.5-4.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Blocked returns a string representing an "blocked" Octicon.
func Blocked(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-blocked%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.467.22a.75.75 0 01.53-.22h6.006a.75.75 0 01.53.22l4.247 4.247c.141.14.22.331.22.53v6.006a.75.75 0 01-.22.53l-4.247 4.247a.75.75 0 01-.53.22H4.997a.75.75 0 01-.53-.22L.22 11.533a.75.75 0 01-.22-.53V4.997a.75.75 0 01.22-.53L4.467.22zm.84 1.28L1.5 5.308v5.384L5.308 14.5h5.384l3.808-3.808V5.308L10.692 1.5H5.308zM4 7.75A.75.75 0 014.75 7h6.5a.75.75 0 010 1.5h-6.5A.75.75 0 014 7.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-blocked%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.638 2.22a.75.75 0 01.53-.22h7.664a.75.75 0 01.53.22l5.418 5.418c.141.14.22.332.22.53v7.664a.75.75 0 01-.22.53l-5.418 5.418a.75.75 0 01-.53.22H8.168a.75.75 0 01-.53-.22l-5.42-5.418a.75.75 0 01-.219-.53V8.168a.75.75 0 01.22-.53l5.418-5.42zM8.48 3.5L3.5 8.48v7.04l4.98 4.98h7.04l4.98-4.98V8.48L15.52 3.5H8.48zM7 11.75a.75.75 0 01.75-.75h8.5a.75.75 0 010 1.5h-8.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Bold returns a string representing an "bold" Octicon.
func Bold(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-bold%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4 2a1 1 0 00-1 1v10a1 1 0 001 1h5.5a3.5 3.5 0 001.852-6.47A3.5 3.5 0 008.5 2H4zm4.5 5a1.5 1.5 0 100-3H5v3h3.5zM5 9v3h4.5a1.5 1.5 0 000-3H5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bold%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 4.75c0-.69.56-1.25 1.25-1.25h5a4.75 4.75 0 013.888 7.479A5 5 0 0114 20.5H7.25c-.69 0-1.25-.56-1.25-1.25V4.75zM8.5 13v5H14a2.5 2.5 0 000-5H8.5zm0-2.5h3.751A2.25 2.25 0 0012.25 6H8.5v4.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Book returns a string representing an "book" Octicon.
func Book(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-book%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 1.75A.75.75 0 01.75 1h4.253c1.227 0 2.317.59 3 1.501A3.744 3.744 0 0111.006 1h4.245a.75.75 0 01.75.75v10.5a.75.75 0 01-.75.75h-4.507a2.25 2.25 0 00-1.591.659l-.622.621a.75.75 0 01-1.06 0l-.622-.621A2.25 2.25 0 005.258 13H.75a.75.75 0 01-.75-.75V1.75zm8.755 3a2.25 2.25 0 012.25-2.25H14.5v9h-3.757c-.71 0-1.4.201-1.992.572l.004-7.322zm-1.504 7.324l.004-5.073-.002-2.253A2.25 2.25 0 005.003 2.5H1.5v9h3.757a3.75 3.75 0 011.994.574z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-book%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 3.75A.75.75 0 01.75 3h7.497c1.566 0 2.945.8 3.751 2.014A4.496 4.496 0 0115.75 3h7.5a.75.75 0 01.75.75v15.063a.75.75 0 01-.755.75l-7.682-.052a3 3 0 00-2.142.878l-.89.891a.75.75 0 01-1.061 0l-.902-.901a3 3 0 00-2.121-.879H.75a.75.75 0 01-.75-.75v-15zm11.247 3.747a3 3 0 00-3-2.997H1.5V18h6.947a4.5 4.5 0 012.803.98l-.003-11.483zm1.503 11.485V7.5a3 3 0 013-3h6.75v13.558l-6.927-.047a4.5 4.5 0 00-2.823.971z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Bookmark returns a string representing an "bookmark" Octicon.
func Bookmark(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-bookmark%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 2.5a.25.25 0 00-.25.25v9.91l3.023-2.489a.75.75 0 01.954 0l3.023 2.49V2.75a.25.25 0 00-.25-.25h-6.5zM3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v11.5a.75.75 0 01-1.227.579L8 11.722l-3.773 3.107A.75.75 0 013 14.25V2.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bookmark%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5 3.75C5 2.784 5.784 2 6.75 2h10.5c.966 0 1.75.784 1.75 1.75v17.5a.75.75 0 01-1.218.586L12 17.21l-5.781 4.625A.75.75 0 015 21.25V3.75zm1.75-.25a.25.25 0 00-.25.25v15.94l5.031-4.026a.75.75 0 01.938 0L17.5 19.69V3.75a.25.25 0 00-.25-.25H6.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// BookmarkFill returns a string representing an "bookmark-fill" Octicon.
func BookmarkFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bookmark-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.69 2a1.75 1.75 0 00-1.75 1.756L5 21.253a.75.75 0 001.219.583L12 17.21l5.782 4.625A.75.75 0 0019 21.25V3.75A1.75 1.75 0 0017.25 2H6.69z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// BookmarkSlash returns a string representing an "bookmark-slash" Octicon.
func BookmarkSlash(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-bookmark-slash%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.19 1.143a.75.75 0 10-.88 1.214L3 4.305v9.945a.75.75 0 001.206.596L8 11.944l3.794 2.902A.75.75 0 0013 14.25v-2.703l1.81 1.31a.75.75 0 10.88-1.214l-2.994-2.168a1.09 1.09 0 00-.014-.01L4.196 3.32a.712.712 0 00-.014-.01L1.19 1.143zM4.5 5.39v7.341l3.044-2.328a.75.75 0 01.912 0l3.044 2.328V10.46l-7-5.07zM5.865 1a.75.75 0 000 1.5h5.385a.25.25 0 01.25.25v3.624a.75.75 0 001.5 0V2.75A1.75 1.75 0 0011.25 1H5.865z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bookmark-slash%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.565 2.018a.75.75 0 00-.88 1.214L5 6.357v14.902a.75.75 0 001.219.585L12 17.21l5.781 4.633A.75.75 0 0019 21.259v-4.764l3.435 2.487a.75.75 0 10.88-1.215L1.565 2.017zM17.5 15.408l-11-7.965v12.254l5.031-4.032a.75.75 0 01.938 0l5.031 4.032v-4.288z\"></path><path d=\"M7.25 2a.75.75 0 000 1.5h10a.25.25 0 01.25.25v6.5a.75.75 0 001.5 0v-6.5A1.75 1.75 0 0017.25 2h-10z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// BookmarkSlashFill returns a string representing an "bookmark-slash-fill" Octicon.
func BookmarkSlashFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bookmark-slash-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.232 2.175a.75.75 0 10-.964 1.15l2.679 2.244L5 21.253a.75.75 0 001.219.583L12 17.21l5.782 4.625A.75.75 0 0019 21.25v-3.907l1.768 1.482a.75.75 0 10.964-1.15l-18.5-15.5zM7.422 2a.75.75 0 00-.482 1.325l10.828 9.073A.75.75 0 0019 11.823V3.75A1.75 1.75 0 0017.25 2H7.421h.001z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Briefcase returns a string representing an "briefcase" Octicon.
func Briefcase(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-briefcase%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.75 0A1.75 1.75 0 005 1.75V3H1.75A1.75 1.75 0 000 4.75v8.5C0 14.216.784 15 1.75 15h12.5A1.75 1.75 0 0016 13.25v-8.5A1.75 1.75 0 0014.25 3H11V1.75A1.75 1.75 0 009.25 0h-2.5zM9.5 3V1.75a.25.25 0 00-.25-.25h-2.5a.25.25 0 00-.25.25V3h3zM5 4.5H1.75a.25.25 0 00-.25.25V6a2 2 0 002 2h9a2 2 0 002-2V4.75a.25.25 0 00-.25-.25H5zm-1.5 5a3.484 3.484 0 01-2-.627v4.377c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V8.873a3.484 3.484 0 01-2 .627h-9z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-briefcase%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.5 1.75C7.5.784 8.284 0 9.25 0h5.5c.966 0 1.75.784 1.75 1.75V4h4.75c.966 0 1.75.784 1.75 1.75v14.5A1.75 1.75 0 0121.25 22H2.75A1.75 1.75 0 011 20.25V5.75C1 4.784 1.784 4 2.75 4H7.5V1.75zm-5 10.24v8.26c0 .138.112.25.25.25h18.5a.25.25 0 00.25-.25v-8.26A4.233 4.233 0 0118.75 13H5.25a4.233 4.233 0 01-2.75-1.01zm19-3.24a2.75 2.75 0 01-2.75 2.75H5.25A2.75 2.75 0 012.5 8.75v-3a.25.25 0 01.25-.25h18.5a.25.25 0 01.25.25v3zm-6.5-7V4H9V1.75a.25.25 0 01.25-.25h5.5a.25.25 0 01.25.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Broadcast returns a string representing an "broadcast" Octicon.
func Broadcast(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-broadcast%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.267 1.457c.3.286.312.76.026 1.06A6.475 6.475 0 001.5 7a6.472 6.472 0 001.793 4.483.75.75 0 01-1.086 1.034 8.89 8.89 0 01-.276-.304l.569-.49-.569.49A7.971 7.971 0 010 7c0-2.139.84-4.083 2.207-5.517a.75.75 0 011.06-.026zm9.466 0a.75.75 0 011.06.026A7.975 7.975 0 0116 7c0 2.139-.84 4.083-2.207 5.517a.75.75 0 11-1.086-1.034A6.475 6.475 0 0014.5 7a6.475 6.475 0 00-1.793-4.483.75.75 0 01.026-1.06zM8.75 8.582a1.75 1.75 0 10-1.5 0v5.668a.75.75 0 001.5 0V8.582zM5.331 4.736a.75.75 0 10-1.143-.972A4.983 4.983 0 003 7c0 1.227.443 2.352 1.177 3.222a.75.75 0 001.146-.967A3.483 3.483 0 014.5 7c0-.864.312-1.654.831-2.264zm6.492-.958a.75.75 0 00-1.146.967c.514.61.823 1.395.823 2.255 0 .86-.31 1.646-.823 2.255a.75.75 0 101.146.967A4.983 4.983 0 0013 7a4.983 4.983 0 00-1.177-3.222z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-broadcast%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M20.485 2.515a.75.75 0 00-1.06 1.06A10.465 10.465 0 0122.5 11c0 2.9-1.174 5.523-3.075 7.424a.75.75 0 001.06 1.061A11.965 11.965 0 0024 11c0-3.314-1.344-6.315-3.515-8.485zm-15.91 1.06a.75.75 0 00-1.06-1.06A11.965 11.965 0 000 11c0 3.313 1.344 6.314 3.515 8.485a.75.75 0 001.06-1.06A10.465 10.465 0 011.5 11c0-2.9 1.174-5.524 3.075-7.425zM8.11 7.11a.75.75 0 00-1.06-1.06A6.98 6.98 0 005 11a6.98 6.98 0 002.05 4.95.75.75 0 001.06-1.061 5.48 5.48 0 01-1.61-3.89 5.48 5.48 0 011.61-3.888zm8.84-1.06a.75.75 0 10-1.06 1.06A5.48 5.48 0 0117.5 11a5.48 5.48 0 01-1.61 3.889.75.75 0 101.06 1.06A6.98 6.98 0 0019 11a6.98 6.98 0 00-2.05-4.949zM14 11a2 2 0 01-1.25 1.855v8.395a.75.75 0 01-1.5 0v-8.395A2 2 0 1114 11z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Browser returns a string representing an "browser" Octicon.
func Browser(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-browser%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 2.75C0 1.784.784 1 1.75 1h12.5c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0114.25 15H1.75A1.75 1.75 0 010 13.25V2.75zm1.75-.25a.25.25 0 00-.25.25V4.5h2v-2H1.75zM5 2.5v2h2v-2H5zm3.5 0v2h6V2.75a.25.25 0 00-.25-.25H8.5zm6 3.5h-13v7.25c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V6z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-browser%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 3.75C0 2.784.784 2 1.75 2h20.5c.966 0 1.75.784 1.75 1.75v16.5A1.75 1.75 0 0122.25 22H1.75A1.75 1.75 0 010 20.25V3.75zm1.75-.25a.25.25 0 00-.25.25V5.5h4v-2H1.75zM7 3.5v2h4v-2H7zm5.5 0v2h10V3.75a.25.25 0 00-.25-.25H12.5zm10 3.5h-21v13.25c0 .138.112.25.25.25h20.5a.25.25 0 00.25-.25V7z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Bug returns a string representing an "bug" Octicon.
func Bug(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-bug%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.72.22a.75.75 0 011.06 0l1 .999a3.492 3.492 0 012.441 0l.999-1a.75.75 0 111.06 1.061l-.775.776c.616.63.995 1.493.995 2.444v.327c0 .1-.009.197-.025.292.408.14.764.392 1.029.722l1.968-.787a.75.75 0 01.556 1.392L13 7.258V9h2.25a.75.75 0 010 1.5H13v.5c0 .409-.049.806-.141 1.186l2.17.868a.75.75 0 01-.557 1.392l-2.184-.873A4.997 4.997 0 018 16a4.997 4.997 0 01-4.288-2.427l-2.183.873a.75.75 0 01-.558-1.392l2.17-.868A5.013 5.013 0 013 11v-.5H.75a.75.75 0 010-1.5H3V7.258L.971 6.446a.75.75 0 01.558-1.392l1.967.787c.265-.33.62-.583 1.03-.722a1.684 1.684 0 01-.026-.292V4.5c0-.951.38-1.814.995-2.444L4.72 1.28a.75.75 0 010-1.06zM6.173 5h3.654A.173.173 0 0010 4.827V4.5a2 2 0 10-4 0v.327c0 .096.077.173.173.173zM5.25 6.5a.75.75 0 00-.75.75V11a3.5 3.5 0 107 0V7.25a.75.75 0 00-.75-.75h-5.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-bug%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.72.22a.75.75 0 011.06 0l1.204 1.203A4.983 4.983 0 0112 1c.717 0 1.4.151 2.016.423L15.22.22a.75.75 0 011.06 1.06l-.971.972A4.988 4.988 0 0117 6v1.104a2.755 2.755 0 011.917 1.974l1.998-.999a.75.75 0 01.67 1.342L19 10.714V13.5l3.25.003a.75.75 0 110 1.5L19 15.001V16a7.02 7.02 0 01-.204 1.686.833.833 0 01.04.018l2.75 1.375a.75.75 0 11-.671 1.342l-2.638-1.319A7 7 0 0112 23a7 7 0 01-6.197-3.742l-2.758 1.181a.75.75 0 11-.59-1.378l2.795-1.199A7.007 7.007 0 015 16v-.996H1.75a.75.75 0 010-1.5H5v-2.79L2.415 9.42a.75.75 0 01.67-1.342l1.998.999A2.755 2.755 0 017 7.104V6a4.99 4.99 0 011.69-3.748l-.97-.972a.75.75 0 010-1.06zM8.5 7h7V6a3.5 3.5 0 10-7 0v1zm-2 3.266V9.75c0-.69.56-1.25 1.25-1.25h8.5c.69 0 1.25.56 1.25 1.25V16a5.5 5.5 0 01-11 0v-5.734z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Calendar returns a string representing an "calendar" Octicon.
func Calendar(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-calendar%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 0a.75.75 0 01.75.75V2h5V.75a.75.75 0 011.5 0V2h1.25c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0113.25 16H2.75A1.75 1.75 0 011 14.25V3.75C1 2.784 1.784 2 2.75 2H4V.75A.75.75 0 014.75 0zm0 3.5h8.5a.25.25 0 01.25.25V6h-11V3.75a.25.25 0 01.25-.25h2zm-2.25 4v6.75c0 .138.112.25.25.25h10.5a.25.25 0 00.25-.25V7.5h-11z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-calendar%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.75 0a.75.75 0 01.75.75V3h9V.75a.75.75 0 011.5 0V3h2.75c.966 0 1.75.784 1.75 1.75v16a1.75 1.75 0 01-1.75 1.75H3.25a1.75 1.75 0 01-1.75-1.75v-16C1.5 3.784 2.284 3 3.25 3H6V.75A.75.75 0 016.75 0zm-3.5 4.5a.25.25 0 00-.25.25V8h18V4.75a.25.25 0 00-.25-.25H3.25zM21 9.5H3v11.25c0 .138.112.25.25.25h17.5a.25.25 0 00.25-.25V9.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Check returns a string representing an "check" Octicon.
func Check(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-check%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13.78 4.22a.75.75 0 010 1.06l-7.25 7.25a.75.75 0 01-1.06 0L2.22 9.28a.75.75 0 011.06-1.06L6 10.94l6.72-6.72a.75.75 0 011.06 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-check%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M21.03 5.72a.75.75 0 010 1.06l-11.5 11.5a.75.75 0 01-1.072-.012l-5.5-5.75a.75.75 0 111.084-1.036l4.97 5.195L19.97 5.72a.75.75 0 011.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CheckCircle returns a string representing an "check-circle" Octicon.
func CheckCircle(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-check-circle%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 8a6.5 6.5 0 1113 0 6.5 6.5 0 01-13 0zM0 8a8 8 0 1116 0A8 8 0 010 8zm11.78-1.72a.75.75 0 00-1.06-1.06L6.75 9.19 5.28 7.72a.75.75 0 00-1.06 1.06l2 2a.75.75 0 001.06 0l4.5-4.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-check-circle%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M17.28 9.28a.75.75 0 00-1.06-1.06l-5.97 5.97-2.47-2.47a.75.75 0 00-1.06 1.06l3 3a.75.75 0 001.06 0l6.5-6.5z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CheckCircleFill returns a string representing an "check-circle-fill" Octicon.
func CheckCircleFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 12:
		return fmt.Sprintf("<svg class=\"octicon octicon-check-circle-fill%s\" height=\"12\" width=\"12\" viewbox=\"0 0 12 12\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 0a6 6 0 100 12A6 6 0 006 0zm-.705 8.737L9.63 4.403 8.392 3.166 5.295 6.263l-1.7-1.702L2.356 5.8l2.938 2.938z\"></path></svg>", extraClasses), true
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-check-circle-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zm3.78-9.72a.75.75 0 00-1.06-1.06L6.75 9.19 5.28 7.72a.75.75 0 00-1.06 1.06l2 2a.75.75 0 001.06 0l4.5-4.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-check-circle-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11-4.925 11-11 11S1 18.075 1 12zm16.28-2.72a.75.75 0 00-1.06-1.06l-5.97 5.97-2.47-2.47a.75.75 0 00-1.06 1.06l3 3a.75.75 0 001.06 0l6.5-6.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Checklist returns a string representing an "checklist" Octicon.
func Checklist(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-checklist%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.5 1.75a.25.25 0 01.25-.25h8.5a.25.25 0 01.25.25v7.736a.75.75 0 101.5 0V1.75A1.75 1.75 0 0011.25 0h-8.5A1.75 1.75 0 001 1.75v11.5c0 .966.784 1.75 1.75 1.75h3.17a.75.75 0 000-1.5H2.75a.25.25 0 01-.25-.25V1.75zM4.75 4a.75.75 0 000 1.5h4.5a.75.75 0 000-1.5h-4.5zM4 7.75A.75.75 0 014.75 7h2a.75.75 0 010 1.5h-2A.75.75 0 014 7.75zm11.774 3.537a.75.75 0 00-1.048-1.074L10.7 14.145 9.281 12.72a.75.75 0 00-1.062 1.058l1.943 1.95a.75.75 0 001.055.008l4.557-4.45z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-checklist%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M3.5 3.75a.25.25 0 01.25-.25h13.5a.25.25 0 01.25.25v10a.75.75 0 001.5 0v-10A1.75 1.75 0 0017.25 2H3.75A1.75 1.75 0 002 3.75v16.5c0 .966.784 1.75 1.75 1.75h7a.75.75 0 000-1.5h-7a.25.25 0 01-.25-.25V3.75z\"></path><path d=\"M6.25 7a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5zm-.75 4.75a.75.75 0 01.75-.75h4.5a.75.75 0 010 1.5h-4.5a.75.75 0 01-.75-.75zm16.28 4.53a.75.75 0 10-1.06-1.06l-4.97 4.97-1.97-1.97a.75.75 0 10-1.06 1.06l2.5 2.5a.75.75 0 001.06 0l5.5-5.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ChevronDown returns a string representing an "chevron-down" Octicon.
func ChevronDown(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-down%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.78 6.22a.75.75 0 010 1.06l-4.25 4.25a.75.75 0 01-1.06 0L3.22 7.28a.75.75 0 011.06-1.06L8 9.94l3.72-3.72a.75.75 0 011.06 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-down%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.22 8.72a.75.75 0 000 1.06l6.25 6.25a.75.75 0 001.06 0l6.25-6.25a.75.75 0 00-1.06-1.06L12 14.44 6.28 8.72a.75.75 0 00-1.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ChevronLeft returns a string representing an "chevron-left" Octicon.
func ChevronLeft(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-left%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.78 12.78a.75.75 0 01-1.06 0L4.47 8.53a.75.75 0 010-1.06l4.25-4.25a.75.75 0 011.06 1.06L6.06 8l3.72 3.72a.75.75 0 010 1.06z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-left%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M15.28 5.22a.75.75 0 00-1.06 0l-6.25 6.25a.75.75 0 000 1.06l6.25 6.25a.75.75 0 101.06-1.06L9.56 12l5.72-5.72a.75.75 0 000-1.06z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ChevronRight returns a string representing an "chevron-right" Octicon.
func ChevronRight(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-right%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.22 3.22a.75.75 0 011.06 0l4.25 4.25a.75.75 0 010 1.06l-4.25 4.25a.75.75 0 01-1.06-1.06L9.94 8 6.22 4.28a.75.75 0 010-1.06z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-right%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.72 18.78a.75.75 0 001.06 0l6.25-6.25a.75.75 0 000-1.06L9.78 5.22a.75.75 0 00-1.06 1.06L14.44 12l-5.72 5.72a.75.75 0 000 1.06z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ChevronUp returns a string representing an "chevron-up" Octicon.
func ChevronUp(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-up%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.22 9.78a.75.75 0 010-1.06l4.25-4.25a.75.75 0 011.06 0l4.25 4.25a.75.75 0 01-1.06 1.06L8 6.06 4.28 9.78a.75.75 0 01-1.06 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-chevron-up%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M18.78 15.28a.75.75 0 000-1.06l-6.25-6.25a.75.75 0 00-1.06 0l-6.25 6.25a.75.75 0 101.06 1.06L12 9.56l5.72 5.72a.75.75 0 001.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Circle returns a string representing an "circle" Octicon.
func Circle(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-circle%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 1.5a6.5 6.5 0 100 13 6.5 6.5 0 000-13zM0 8a8 8 0 1116 0A8 8 0 010 8z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-circle%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 2.5a9.5 9.5 0 100 19 9.5 9.5 0 000-19zM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11-4.925 11-11 11S1 18.075 1 12z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CircleSlash returns a string representing an "circle-slash" Octicon.
func CircleSlash(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-circle-slash%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 8a6.5 6.5 0 0110.535-5.096l-9.131 9.131A6.472 6.472 0 011.5 8zm2.465 5.096a6.5 6.5 0 009.131-9.131l-9.131 9.131zM8 0a8 8 0 100 16A8 8 0 008 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-circle-slash%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12A9.5 9.5 0 0112 2.5c2.353 0 4.507.856 6.166 2.273L4.773 18.166A9.462 9.462 0 012.5 12zm3.334 7.227A9.462 9.462 0 0012 21.5a9.5 9.5 0 009.5-9.5 9.462 9.462 0 00-2.273-6.166L5.834 19.227z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Clock returns a string representing an "clock" Octicon.
func Clock(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-clock%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 8a6.5 6.5 0 1113 0 6.5 6.5 0 01-13 0zM8 0a8 8 0 100 16A8 8 0 008 0zm.5 4.75a.75.75 0 00-1.5 0v3.5a.75.75 0 00.471.696l2.5 1a.75.75 0 00.557-1.392L8.5 7.742V4.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-clock%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12.5 7.25a.75.75 0 00-1.5 0v5.5c0 .27.144.518.378.651l3.5 2a.75.75 0 00.744-1.302L12.5 12.315V7.25z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Cloud returns a string representing an "cloud" Octicon.
func Cloud(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-cloud%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 7.25A5.225 5.225 0 017.25 2a5.222 5.222 0 014.767 3.029A4.472 4.472 0 0116 9.5c0 2.505-1.995 4.5-4.5 4.5h-8A3.475 3.475 0 010 10.5c0-1.41.809-2.614 2.001-3.17L2 7.25zm1.54.482a.75.75 0 01-.556.832c-.86.22-1.484.987-1.484 1.936 0 1.124.876 2 2 2h8c1.676 0 3-1.324 3-3s-1.324-3-3-3a.75.75 0 01-.709-.504A3.72 3.72 0 007.25 3.5C5.16 3.5 3.5 5.16 3.5 7.25a3.276 3.276 0 00.035.436l.004.036.001.008v.002z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-cloud%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.103 10.107c0-4.244 3.445-7.607 7.733-7.607 3.19 0 5.912 1.858 7.099 4.563a.634.634 0 01.01.022l.001.006C21.348 7.345 24 10.095 24 13.536 24 17.148 21.076 20 17.431 20H5.017C2.23 20 0 17.83 0 15.06a4.9 4.9 0 013.112-4.581 8.024 8.024 0 01-.009-.372zM10.836 4c-3.485 0-6.233 2.717-6.233 6.107 0 .284.022.602.052.756a.75.75 0 01-.552.869c-1.52.385-2.603 1.712-2.603 3.328 0 1.917 1.532 3.44 3.517 3.44h12.414c2.843 0 5.069-2.206 5.069-4.964 0-2.759-2.226-4.965-5.069-4.965a.75.75 0 01-.696-.47l-.179-.446C15.606 5.5 13.424 4 10.836 4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CloudOffline returns a string representing an "cloud-offline" Octicon.
func CloudOffline(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-cloud-offline%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M7.25 2c-.69 0-1.351.13-1.957.371a.75.75 0 10.554 1.394c.43-.17.903-.265 1.403-.265a3.72 3.72 0 013.541 2.496.75.75 0 00.709.504c1.676 0 3 1.324 3 3a3 3 0 01-.681 1.92.75.75 0 001.156.955A4.496 4.496 0 0016 9.5a4.472 4.472 0 00-3.983-4.471A5.222 5.222 0 007.25 2z\"></path><path fill-rule=\"evenodd\" d=\"M.72 1.72a.75.75 0 011.06 0l2.311 2.31c.03.025.056.052.08.08l8.531 8.532a.785.785 0 01.035.034l2.043 2.044a.75.75 0 11-1.06 1.06l-1.8-1.799a4.64 4.64 0 01-.42.019h-8A3.475 3.475 0 010 10.5c0-1.41.809-2.614 2.001-3.17a5.218 5.218 0 01.646-2.622L.72 2.78a.75.75 0 010-1.06zM3.5 7.25c0-.505.096-.983.271-1.418L10.44 12.5H3.5c-1.124 0-2-.876-2-2 0-.95.624-1.716 1.484-1.936a.75.75 0 00.557-.833A4.1 4.1 0 013.5 7.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-cloud-offline%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.78 2.22a.75.75 0 00-1.06 1.06l2.64 2.642a7.525 7.525 0 00-1.249 4.557A4.9 4.9 0 000 15.059C0 17.831 2.229 20 5.017 20h12.414a6.879 6.879 0 00.944-.065l2.845 2.845a.75.75 0 101.06-1.06L2.78 2.22zM16.94 18.5L5.448 7.01a6.03 6.03 0 00-.794 3.853.75.75 0 01-.552.869c-1.52.385-2.603 1.712-2.603 3.328 0 1.917 1.532 3.44 3.517 3.44H16.94z\"></path><path d=\"M10.836 2.5a7.865 7.865 0 00-3.638.88.75.75 0 10.692 1.331A6.365 6.365 0 0110.836 4c2.588 0 4.77 1.5 5.72 3.655l.179.445a.75.75 0 00.696.471c2.843 0 5.069 2.206 5.069 4.965a4.9 4.9 0 01-1.684 3.716.75.75 0 00.986 1.13A6.396 6.396 0 0024 13.536c0-3.44-2.652-6.191-6.054-6.445l-.002-.006a.634.634 0 00-.01-.022C16.749 4.358 14.026 2.5 10.837 2.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Code returns a string representing an "code" Octicon.
func Code(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-code%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.72 3.22a.75.75 0 011.06 1.06L2.06 8l3.72 3.72a.75.75 0 11-1.06 1.06L.47 8.53a.75.75 0 010-1.06l4.25-4.25zm6.56 0a.75.75 0 10-1.06 1.06L13.94 8l-3.72 3.72a.75.75 0 101.06 1.06l4.25-4.25a.75.75 0 000-1.06l-4.25-4.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-code%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.78 4.97a.75.75 0 010 1.06L2.81 12l5.97 5.97a.75.75 0 11-1.06 1.06l-6.5-6.5a.75.75 0 010-1.06l6.5-6.5a.75.75 0 011.06 0zm6.44 0a.75.75 0 000 1.06L21.19 12l-5.97 5.97a.75.75 0 101.06 1.06l6.5-6.5a.75.75 0 000-1.06l-6.5-6.5a.75.75 0 00-1.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CodeOfConduct returns a string representing an "code-of-conduct" Octicon.
func CodeOfConduct(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-code-of-conduct%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.048 2.241c.964-.709 2.079-1.238 3.325-1.241a4.613 4.613 0 013.282 1.355c.41.408.757.86.996 1.428.238.568.348 1.206.347 1.968 0 2.193-1.505 4.254-3.081 5.862-1.496 1.526-3.213 2.796-4.249 3.563l-.22.163a.75.75 0 01-.895 0l-.221-.163c-1.036-.767-2.753-2.037-4.249-3.563C1.51 10.008.007 7.952.002 5.762a4.614 4.614 0 011.353-3.407C3.123.585 6.223.537 8.048 2.24zm-1.153.983c-.81.78-1.546 1.669-2.166 2.417-.184.222-.358.432-.52.623a.75.75 0 00.04 1.016c.35.35.697.697 1.043 1.047.866.875 2.292.914 3.185.032.264-.26.534-.528.802-.797.694-.694 1.8-.701 2.474-.03L12.92 8.7l.283.284c-.244.334-.515.666-.81.995l-1.384-1.28A.75.75 0 109.99 9.802l1.357 1.252c-.325.31-.656.606-.984.887l-1.48-1.366a.75.75 0 10-1.018 1.102L9.191 12.9c-.433.34-.838.643-1.191.905-1.04-.773-2.537-1.907-3.846-3.242C2.611 8.99 1.502 7.306 1.502 5.75a3.114 3.114 0 01.913-2.335c1.159-1.158 3.23-1.224 4.48-.191zm7.112 4.442c.313-.65.491-1.293.491-1.916v-.001c0-.614-.088-1.045-.23-1.385-.143-.339-.357-.633-.673-.949a3.113 3.113 0 00-2.218-.915c-1.092.003-2.165.627-3.226 1.602-.823.755-1.554 1.637-2.228 2.45l-.127.154.562.566a.756.756 0 001.066.02l.794-.79c1.258-1.258 3.312-1.31 4.594-.032.396.394.792.791 1.173 1.173l.022.023z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-code-of-conduct%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.828 4.328C5.26 1.896 9.5 1.881 11.935 4.317c.024.024.046.05.067.076 1.391-1.078 2.993-1.886 4.777-1.89a6.216 6.216 0 014.424 1.825c.559.56 1.023 1.165 1.34 1.922.318.756.47 1.617.468 2.663 0 2.972-2.047 5.808-4.269 8.074-2.098 2.14-4.507 3.924-5.974 5.009l-.311.23a.752.752 0 01-.897 0l-.312-.23c-1.466-1.085-3.875-2.869-5.973-5.009-2.22-2.263-4.264-5.095-4.27-8.063v.012-.024.012a6.217 6.217 0 011.823-4.596zm8.033 1.042c-1.846-1.834-5.124-1.823-6.969.022a4.713 4.713 0 00-1.382 3.52c0 2.332 1.65 4.79 3.839 7.022 1.947 1.986 4.184 3.66 5.66 4.752a79.983 79.983 0 002.159-1.645l-2.14-1.974a.752.752 0 011.02-1.106l2.295 2.118c.616-.52 1.242-1.08 1.85-1.672l-2.16-1.992a.752.752 0 011.021-1.106l2.188 2.02a18.992 18.992 0 001.528-1.877l-.585-.586-1.651-1.652c-1.078-1.074-2.837-1.055-3.935.043-.379.38-.76.758-1.132 1.126-1.14 1.124-2.96 1.077-4.07-.043-.489-.495-.98-.988-1.475-1.482a.752.752 0 01-.04-1.019c.234-.276.483-.576.745-.893.928-1.12 2.023-2.442 3.234-3.576zm9.725 6.77c.579-1.08.92-2.167.92-3.228.002-.899-.128-1.552-.35-2.08-.22-.526-.551-.974-1.017-1.44a4.71 4.71 0 00-3.356-1.384c-1.66.004-3.25.951-4.77 2.346-1.18 1.084-2.233 2.353-3.188 3.506l-.351.423c.331.332.663.664.993.998a1.375 1.375 0 001.943.03c.37-.365.748-.74 1.125-1.118 1.662-1.663 4.373-1.726 6.06-.045.56.558 1.12 1.12 1.658 1.658l.333.334z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CodeReview returns a string representing an "code-review" Octicon.
func CodeReview(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-code-review%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 2.75a.25.25 0 01.25-.25h12.5a.25.25 0 01.25.25v8.5a.25.25 0 01-.25.25h-6.5a.75.75 0 00-.53.22L4.5 14.44v-2.19a.75.75 0 00-.75-.75h-2a.25.25 0 01-.25-.25v-8.5zM1.75 1A1.75 1.75 0 000 2.75v8.5C0 12.216.784 13 1.75 13H3v1.543a1.457 1.457 0 002.487 1.03L8.061 13h6.189A1.75 1.75 0 0016 11.25v-8.5A1.75 1.75 0 0014.25 1H1.75zm5.03 3.47a.75.75 0 010 1.06L5.31 7l1.47 1.47a.75.75 0 01-1.06 1.06l-2-2a.75.75 0 010-1.06l2-2a.75.75 0 011.06 0zm2.44 0a.75.75 0 000 1.06L10.69 7 9.22 8.47a.75.75 0 001.06 1.06l2-2a.75.75 0 000-1.06l-2-2a.75.75 0 00-1.06 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-code-review%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M10.3 6.74a.75.75 0 01-.04 1.06l-2.908 2.7 2.908 2.7a.75.75 0 11-1.02 1.1l-3.5-3.25a.75.75 0 010-1.1l3.5-3.25a.75.75 0 011.06.04zm3.44 1.06a.75.75 0 111.02-1.1l3.5 3.25a.75.75 0 010 1.1l-3.5 3.25a.75.75 0 11-1.02-1.1l2.908-2.7-2.908-2.7z\"></path><path fill-rule=\"evenodd\" d=\"M1.5 4.25c0-.966.784-1.75 1.75-1.75h17.5c.966 0 1.75.784 1.75 1.75v12.5a1.75 1.75 0 01-1.75 1.75h-9.69l-3.573 3.573A1.457 1.457 0 015 21.043V18.5H3.25a1.75 1.75 0 01-1.75-1.75V4.25zM3.25 4a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h2.5a.75.75 0 01.75.75v3.19l3.72-3.72a.75.75 0 01.53-.22h10a.25.25 0 00.25-.25V4.25a.25.25 0 00-.25-.25H3.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CodeSquare returns a string representing an "code-square" Octicon.
func CodeSquare(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-code-square%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 1.5a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V1.75a.25.25 0 00-.25-.25H1.75zM0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v12.5A1.75 1.75 0 0114.25 16H1.75A1.75 1.75 0 010 14.25V1.75zm9.22 3.72a.75.75 0 000 1.06L10.69 8 9.22 9.47a.75.75 0 101.06 1.06l2-2a.75.75 0 000-1.06l-2-2a.75.75 0 00-1.06 0zM6.78 6.53a.75.75 0 00-1.06-1.06l-2 2a.75.75 0 000 1.06l2 2a.75.75 0 101.06-1.06L5.31 8l1.47-1.47z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-code-square%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M10.3 8.24a.75.75 0 01-.04 1.06L7.352 12l2.908 2.7a.75.75 0 11-1.02 1.1l-3.5-3.25a.75.75 0 010-1.1l3.5-3.25a.75.75 0 011.06.04zm3.44 1.06a.75.75 0 111.02-1.1l3.5 3.25a.75.75 0 010 1.1l-3.5 3.25a.75.75 0 11-1.02-1.1l2.908-2.7-2.908-2.7z\"></path><path fill-rule=\"evenodd\" d=\"M2 3.75C2 2.784 2.784 2 3.75 2h16.5c.966 0 1.75.784 1.75 1.75v16.5A1.75 1.75 0 0120.25 22H3.75A1.75 1.75 0 012 20.25V3.75zm1.75-.25a.25.25 0 00-.25.25v16.5c0 .138.112.25.25.25h16.5a.25.25 0 00.25-.25V3.75a.25.25 0 00-.25-.25H3.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Codescan returns a string representing an "codescan" Octicon.
func Codescan(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-codescan%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8.47 4.97a.75.75 0 000 1.06L9.94 7.5 8.47 8.97a.75.75 0 101.06 1.06l2-2a.75.75 0 000-1.06l-2-2a.75.75 0 00-1.06 0zM6.53 6.03a.75.75 0 00-1.06-1.06l-2 2a.75.75 0 000 1.06l2 2a.75.75 0 101.06-1.06L5.06 7.5l1.47-1.47z\"></path><path fill-rule=\"evenodd\" d=\"M12.246 13.307a7.5 7.5 0 111.06-1.06l2.474 2.473a.75.75 0 11-1.06 1.06l-2.474-2.473zM1.5 7.5a6 6 0 1110.386 4.094.75.75 0 00-.292.293A6 6 0 011.5 7.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-codescan%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M11.97 6.97a.75.75 0 000 1.06l2.47 2.47-2.47 2.47a.75.75 0 101.06 1.06l3-3a.75.75 0 000-1.06l-3-3a.75.75 0 00-1.06 0zM9.03 8.03a.75.75 0 00-1.06-1.06l-3 3a.75.75 0 000 1.06l3 3a.75.75 0 001.06-1.06L6.56 10.5l2.47-2.47z\"></path><path fill-rule=\"evenodd\" d=\"M10.5 0C4.701 0 0 4.701 0 10.5S4.701 21 10.5 21c2.63 0 5.033-.967 6.875-2.564l4.345 4.344a.75.75 0 101.06-1.06l-4.344-4.345A10.459 10.459 0 0021 10.5C21 4.701 16.299 0 10.5 0zm-9 10.5a9 9 0 1118 0 9 9 0 01-18 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CodescanCheckmark returns a string representing an "codescan-checkmark" Octicon.
func CodescanCheckmark(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-codescan-checkmark%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M10.28 6.28a.75.75 0 10-1.06-1.06L6.25 8.19l-.97-.97a.75.75 0 00-1.06 1.06l1.5 1.5a.75.75 0 001.06 0l3.5-3.5z\"></path><path fill-rule=\"evenodd\" d=\"M7.5 15a7.469 7.469 0 004.746-1.693l2.474 2.473a.75.75 0 101.06-1.06l-2.473-2.474A7.5 7.5 0 107.5 15zm0-13.5a6 6 0 104.094 10.386.75.75 0 01.293-.292A6 6 0 007.5 1.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-codescan-checkmark%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M15.03 8.28a.75.75 0 00-1.06-1.06l-5.22 5.22-2.22-2.22a.75.75 0 10-1.06 1.06l2.75 2.75a.75.75 0 001.06 0l5.75-5.75z\"></path><path fill-rule=\"evenodd\" d=\"M0 10.5C0 4.701 4.701 0 10.5 0S21 4.701 21 10.5c0 2.63-.967 5.033-2.564 6.875l4.344 4.345a.75.75 0 11-1.06 1.06l-4.345-4.344A10.459 10.459 0 0110.5 21C4.701 21 0 16.299 0 10.5zm10.5-9a9 9 0 100 18 9 9 0 000-18z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Codespaces returns a string representing an "codespaces" Octicon.
func Codespaces(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-codespaces%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 1.75C2 .784 2.784 0 3.75 0h8.5C13.216 0 14 .784 14 1.75v5a1.75 1.75 0 01-1.75 1.75h-8.5A1.75 1.75 0 012 6.75v-5zm1.75-.25a.25.25 0 00-.25.25v5c0 .138.112.25.25.25h8.5a.25.25 0 00.25-.25v-5a.25.25 0 00-.25-.25h-8.5zM0 11.25c0-.966.784-1.75 1.75-1.75h12.5c.966 0 1.75.784 1.75 1.75v3A1.75 1.75 0 0114.25 16H1.75A1.75 1.75 0 010 14.25v-3zM1.75 11a.25.25 0 00-.25.25v3c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25v-3a.25.25 0 00-.25-.25H1.75z\"></path><path fill-rule=\"evenodd\" d=\"M3 12.75a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5a.75.75 0 01-.75-.75zm4 0a.75.75 0 01.75-.75h4.5a.75.75 0 010 1.5h-4.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-codespaces%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.5 3.75C3.5 2.784 4.284 2 5.25 2h13.5c.966 0 1.75.784 1.75 1.75v7.5A1.75 1.75 0 0118.75 13H5.25a1.75 1.75 0 01-1.75-1.75v-7.5zm1.75-.25a.25.25 0 00-.25.25v7.5c0 .138.112.25.25.25h13.5a.25.25 0 00.25-.25v-7.5a.25.25 0 00-.25-.25H5.25zM1.5 15.75c0-.966.784-1.75 1.75-1.75h17.5c.966 0 1.75.784 1.75 1.75v4a1.75 1.75 0 01-1.75 1.75H3.25a1.75 1.75 0 01-1.75-1.75v-4zm1.75-.25a.25.25 0 00-.25.25v4c0 .138.112.25.25.25h17.5a.25.25 0 00.25-.25v-4a.25.25 0 00-.25-.25H3.25z\"></path><path fill-rule=\"evenodd\" d=\"M10 17.75a.75.75 0 01.75-.75h6.5a.75.75 0 010 1.5h-6.5a.75.75 0 01-.75-.75zm-4 0a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Columns returns a string representing an "columns" Octicon.
func Columns(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-columns%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 0A1.75 1.75 0 001 1.75v12.5c0 .966.784 1.75 1.75 1.75h2.5A1.75 1.75 0 007 14.25V1.75A1.75 1.75 0 005.25 0h-2.5zM2.5 1.75a.25.25 0 01.25-.25h2.5a.25.25 0 01.25.25v12.5a.25.25 0 01-.25.25h-2.5a.25.25 0 01-.25-.25V1.75zM10.75 0A1.75 1.75 0 009 1.75v12.5c0 .966.784 1.75 1.75 1.75h2.5A1.75 1.75 0 0015 14.25V1.75A1.75 1.75 0 0013.25 0h-2.5zm-.25 1.75a.25.25 0 01.25-.25h2.5a.25.25 0 01.25.25v12.5a.25.25 0 01-.25.25h-2.5a.25.25 0 01-.25-.25V1.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-columns%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 2A1.75 1.75 0 002 3.75v16.5c0 .966.784 1.75 1.75 1.75h5.5A1.75 1.75 0 0011 20.25V3.75A1.75 1.75 0 009.25 2h-5.5zM3.5 3.75a.25.25 0 01.25-.25h5.5a.25.25 0 01.25.25v16.5a.25.25 0 01-.25.25h-5.5a.25.25 0 01-.25-.25V3.75zM14.75 2A1.75 1.75 0 0013 3.75v16.5c0 .966.784 1.75 1.75 1.75h5.5A1.75 1.75 0 0022 20.25V3.75A1.75 1.75 0 0020.25 2h-5.5zm-.25 1.75a.25.25 0 01.25-.25h5.5a.25.25 0 01.25.25v16.5a.25.25 0 01-.25.25h-5.5a.25.25 0 01-.25-.25V3.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Comment returns a string representing an "comment" Octicon.
func Comment(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-comment%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 2.5a.25.25 0 00-.25.25v7.5c0 .138.112.25.25.25h2a.75.75 0 01.75.75v2.19l2.72-2.72a.75.75 0 01.53-.22h4.5a.25.25 0 00.25-.25v-7.5a.25.25 0 00-.25-.25H2.75zM1 2.75C1 1.784 1.784 1 2.75 1h10.5c.966 0 1.75.784 1.75 1.75v7.5A1.75 1.75 0 0113.25 12H9.06l-2.573 2.573A1.457 1.457 0 014 13.543V12H2.75A1.75 1.75 0 011 10.25v-7.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-comment%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.25 4a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h2.5a.75.75 0 01.75.75v3.19l3.72-3.72a.75.75 0 01.53-.22h10a.25.25 0 00.25-.25V4.25a.25.25 0 00-.25-.25H3.25zm-1.75.25c0-.966.784-1.75 1.75-1.75h17.5c.966 0 1.75.784 1.75 1.75v12.5a1.75 1.75 0 01-1.75 1.75h-9.69l-3.573 3.573A1.457 1.457 0 015 21.043V18.5H3.25a1.75 1.75 0 01-1.75-1.75V4.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CommentDiscussion returns a string representing an "comment-discussion" Octicon.
func CommentDiscussion(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-comment-discussion%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 2.75a.25.25 0 01.25-.25h8.5a.25.25 0 01.25.25v5.5a.25.25 0 01-.25.25h-3.5a.75.75 0 00-.53.22L3.5 11.44V9.25a.75.75 0 00-.75-.75h-1a.25.25 0 01-.25-.25v-5.5zM1.75 1A1.75 1.75 0 000 2.75v5.5C0 9.216.784 10 1.75 10H2v1.543a1.457 1.457 0 002.487 1.03L7.061 10h3.189A1.75 1.75 0 0012 8.25v-5.5A1.75 1.75 0 0010.25 1h-8.5zM14.5 4.75a.25.25 0 00-.25-.25h-.5a.75.75 0 110-1.5h.5c.966 0 1.75.784 1.75 1.75v5.5A1.75 1.75 0 0114.25 12H14v1.543a1.457 1.457 0 01-2.487 1.03L9.22 12.28a.75.75 0 111.06-1.06l2.22 2.22v-2.19a.75.75 0 01.75-.75h1a.25.25 0 00.25-.25v-5.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-comment-discussion%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 1A1.75 1.75 0 000 2.75v9.5C0 13.216.784 14 1.75 14H3v1.543a1.457 1.457 0 002.487 1.03L8.061 14h6.189A1.75 1.75 0 0016 12.25v-9.5A1.75 1.75 0 0014.25 1H1.75zM1.5 2.75a.25.25 0 01.25-.25h12.5a.25.25 0 01.25.25v9.5a.25.25 0 01-.25.25h-6.5a.75.75 0 00-.53.22L4.5 15.44v-2.19a.75.75 0 00-.75-.75h-2a.25.25 0 01-.25-.25v-9.5z\"></path><path d=\"M22.5 8.75a.25.25 0 00-.25-.25h-3.5a.75.75 0 010-1.5h3.5c.966 0 1.75.784 1.75 1.75v9.5A1.75 1.75 0 0122.25 20H21v1.543a1.457 1.457 0 01-2.487 1.03L15.939 20H10.75A1.75 1.75 0 019 18.25v-1.465a.75.75 0 011.5 0v1.465c0 .138.112.25.25.25h5.5a.75.75 0 01.53.22l2.72 2.72v-2.19a.75.75 0 01.75-.75h2a.25.25 0 00.25-.25v-9.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Commit returns a string representing an "commit" Octicon.
func Commit(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-commit%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M17.5 11.75a.75.75 0 01.75-.75h5a.75.75 0 010 1.5h-5a.75.75 0 01-.75-.75zm-17.5 0A.75.75 0 01.75 11h5a.75.75 0 010 1.5h-5a.75.75 0 01-.75-.75z\"></path><path fill-rule=\"evenodd\" d=\"M12 16.25a4.5 4.5 0 100-9 4.5 4.5 0 000 9zm0 1.5a6 6 0 100-12 6 6 0 000 12z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Container returns a string representing an "container" Octicon.
func Container(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-container%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.41.24l4.711 2.774A1.767 1.767 0 0116 4.54v5.01a1.77 1.77 0 01-.88 1.53l-7.753 4.521-.002.001a1.767 1.767 0 01-1.774 0H5.59L.873 12.85A1.762 1.762 0 010 11.327V6.292c0-.304.078-.598.22-.855l.004-.005.01-.019c.15-.262.369-.486.64-.643L8.641.239a1.75 1.75 0 011.765 0l.002.001zM9.397 1.534a.25.25 0 01.252 0l4.115 2.422-7.152 4.148a.267.267 0 01-.269 0L2.227 5.716l7.17-4.182zM7.365 9.402L8.73 8.61v4.46l-1.5.875V9.473a1.77 1.77 0 00.136-.071zm2.864 2.794V7.741l1.521-.882v4.45l-1.521.887zm3.021-1.762l1.115-.65h.002a.268.268 0 00.133-.232V5.264l-1.25.725v4.445zm-11.621 1.12l4.1 2.393V9.474a1.77 1.77 0 01-.138-.072L1.5 7.029v4.298c0 .095.05.181.129.227z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-container%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13.152.682a2.25 2.25 0 012.269 0l.007.004 6.957 4.276a2.276 2.276 0 011.126 1.964v7.516c0 .81-.432 1.56-1.133 1.968l-.002.001-11.964 7.037-.004.003a2.276 2.276 0 01-2.284 0l-.026-.015-6.503-4.502a2.268 2.268 0 01-1.096-1.943V9.438c0-.392.1-.77.284-1.1l.003-.006.014-.026a2.28 2.28 0 01.82-.827h.002L13.152.681zm.757 1.295h-.001L2.648 8.616l6.248 4.247a.776.776 0 00.758-.01h.001l11.633-6.804-6.629-4.074a.75.75 0 00-.75.003zM18 9.709l-3.25 1.9v7.548L18 17.245V9.709zm1.5-.878v7.532l2.124-1.25a.777.777 0 00.387-.671V7.363L19.5 8.831zm-9.09 5.316l2.84-1.66v7.552l-3.233 1.902v-7.612c.134-.047.265-.107.391-.18l.002-.002zm-1.893 7.754V14.33a2.277 2.277 0 01-.393-.18l-.023-.014-6.102-4.147v7.003c0 .275.145.528.379.664l.025.014 6.114 4.232z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Copy returns a string representing an "copy" Octicon.
func Copy(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-copy%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 010 1.5h-1.5a.25.25 0 00-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 00.25-.25v-1.5a.75.75 0 011.5 0v1.5A1.75 1.75 0 019.25 16h-7.5A1.75 1.75 0 010 14.25v-7.5z\"></path><path fill-rule=\"evenodd\" d=\"M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0114.25 11h-7.5A1.75 1.75 0 015 9.25v-7.5zm1.75-.25a.25.25 0 00-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 00.25-.25v-7.5a.25.25 0 00-.25-.25h-7.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-copy%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.024 3.75c0-.966.784-1.75 1.75-1.75H20.25c.966 0 1.75.784 1.75 1.75v11.498a1.75 1.75 0 01-1.75 1.75H8.774a1.75 1.75 0 01-1.75-1.75V3.75zm1.75-.25a.25.25 0 00-.25.25v11.498c0 .139.112.25.25.25H20.25a.25.25 0 00.25-.25V3.75a.25.25 0 00-.25-.25H8.774z\"></path><path d=\"M1.995 10.749a1.75 1.75 0 011.75-1.751H5.25a.75.75 0 110 1.5H3.745a.25.25 0 00-.25.25L3.5 20.25c0 .138.111.25.25.25h9.5a.25.25 0 00.25-.25v-1.51a.75.75 0 111.5 0v1.51A1.75 1.75 0 0113.25 22h-9.5A1.75 1.75 0 012 20.25l-.005-9.501z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Cpu returns a string representing an "cpu" Octicon.
func Cpu(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-cpu%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.5.75a.75.75 0 00-1.5 0V2H3.75A1.75 1.75 0 002 3.75V5H.75a.75.75 0 000 1.5H2v3H.75a.75.75 0 000 1.5H2v1.25c0 .966.784 1.75 1.75 1.75H5v1.25a.75.75 0 001.5 0V14h3v1.25a.75.75 0 001.5 0V14h1.25A1.75 1.75 0 0014 12.25V11h1.25a.75.75 0 000-1.5H14v-3h1.25a.75.75 0 000-1.5H14V3.75A1.75 1.75 0 0012.25 2H11V.75a.75.75 0 00-1.5 0V2h-3V.75zm5.75 11.75h-8.5a.25.25 0 01-.25-.25v-8.5a.25.25 0 01.25-.25h8.5a.25.25 0 01.25.25v8.5a.25.25 0 01-.25.25zM5.75 5a.75.75 0 00-.75.75v4.5c0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75v-4.5a.75.75 0 00-.75-.75h-4.5zm.75 4.5v-3h3v3h-3z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-cpu%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.75 8a.75.75 0 00-.75.75v6.5c0 .414.336.75.75.75h6.5a.75.75 0 00.75-.75v-6.5a.75.75 0 00-.75-.75h-6.5zm.75 6.5v-5h5v5h-5z\"></path><path fill-rule=\"evenodd\" d=\"M15.25 1a.75.75 0 01.75.75V4h2.25c.966 0 1.75.784 1.75 1.75V8h2.25a.75.75 0 010 1.5H20v5h2.25a.75.75 0 010 1.5H20v2.25A1.75 1.75 0 0118.25 20H16v2.25a.75.75 0 01-1.5 0V20h-5v2.25a.75.75 0 01-1.5 0V20H5.75A1.75 1.75 0 014 18.25V16H1.75a.75.75 0 010-1.5H4v-5H1.75a.75.75 0 010-1.5H4V5.75C4 4.784 4.784 4 5.75 4H8V1.75a.75.75 0 011.5 0V4h5V1.75a.75.75 0 01.75-.75zm3 17.5a.25.25 0 00.25-.25V5.75a.25.25 0 00-.25-.25H5.75a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h12.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CreditCard returns a string representing an "credit-card" Octicon.
func CreditCard(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-credit-card%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M10.75 9a.75.75 0 000 1.5h1.5a.75.75 0 000-1.5h-1.5z\"></path><path fill-rule=\"evenodd\" d=\"M0 3.75C0 2.784.784 2 1.75 2h12.5c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0114.25 14H1.75A1.75 1.75 0 010 12.25v-8.5zm14.5 0V5h-13V3.75a.25.25 0 01.25-.25h12.5a.25.25 0 01.25.25zm0 2.75h-13v5.75c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V6.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-credit-card%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M15.25 14a.75.75 0 000 1.5h3.5a.75.75 0 000-1.5h-3.5z\"></path><path fill-rule=\"evenodd\" d=\"M1.75 3A1.75 1.75 0 000 4.75v14.5C0 20.216.784 21 1.75 21h20.5A1.75 1.75 0 0024 19.25V4.75A1.75 1.75 0 0022.25 3H1.75zM1.5 4.75a.25.25 0 01.25-.25h20.5a.25.25 0 01.25.25V8.5h-21V4.75zm0 5.25v9.25c0 .138.112.25.25.25h20.5a.25.25 0 00.25-.25V10h-21z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// CrossReference returns a string representing an "cross-reference" Octicon.
func CrossReference(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-cross-reference%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M16 1.25v4.146a.25.25 0 01-.427.177L14.03 4.03l-3.75 3.75a.75.75 0 11-1.06-1.06l3.75-3.75-1.543-1.543A.25.25 0 0111.604 1h4.146a.25.25 0 01.25.25zM2.75 3.5a.25.25 0 00-.25.25v7.5c0 .138.112.25.25.25h2a.75.75 0 01.75.75v2.19l2.72-2.72a.75.75 0 01.53-.22h4.5a.25.25 0 00.25-.25v-2.5a.75.75 0 111.5 0v2.5A1.75 1.75 0 0113.25 13H9.06l-2.573 2.573A1.457 1.457 0 014 14.543V13H2.75A1.75 1.75 0 011 11.25v-7.5C1 2.784 1.784 2 2.75 2h5.5a.75.75 0 010 1.5h-5.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-cross-reference%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M16.5 2.25a.75.75 0 01.75-.75h5.5a.75.75 0 01.75.75v5.5a.75.75 0 01-1.5 0V4.06l-6.22 6.22a.75.75 0 11-1.06-1.06L20.94 3h-3.69a.75.75 0 01-.75-.75z\"></path><path d=\"M3.25 4a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h2.5a.75.75 0 01.75.75v3.19l3.72-3.72a.75.75 0 01.53-.22h10a.25.25 0 00.25-.25v-6a.75.75 0 011.5 0v6a1.75 1.75 0 01-1.75 1.75h-9.69l-3.573 3.573A1.457 1.457 0 015 21.043V18.5H3.25a1.75 1.75 0 01-1.75-1.75V4.25c0-.966.784-1.75 1.75-1.75h11a.75.75 0 010 1.5h-11z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Dash returns a string representing an "dash" Octicon.
func Dash(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-dash%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 7.75A.75.75 0 012.75 7h10a.75.75 0 010 1.5h-10A.75.75 0 012 7.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-dash%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.5 12.75a.75.75 0 01.75-.75h13.5a.75.75 0 010 1.5H5.25a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Database returns a string representing an "database" Octicon.
func Database(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-database%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.5 3.5c0-.133.058-.318.282-.55.227-.237.592-.484 1.1-.708C4.899 1.795 6.354 1.5 8 1.5c1.647 0 3.102.295 4.117.742.51.224.874.47 1.101.707.224.233.282.418.282.551 0 .133-.058.318-.282.55-.227.237-.592.484-1.1.708C11.101 5.205 9.646 5.5 8 5.5c-1.647 0-3.102-.295-4.117-.742-.51-.224-.874-.47-1.101-.707-.224-.233-.282-.418-.282-.551zM1 3.5c0-.626.292-1.165.7-1.59.406-.422.956-.767 1.579-1.041C4.525.32 6.195 0 8 0c1.805 0 3.475.32 4.722.869.622.274 1.172.62 1.578 1.04.408.426.7.965.7 1.591v9c0 .626-.292 1.165-.7 1.59-.406.422-.956.767-1.579 1.041C11.476 15.68 9.806 16 8 16c-1.805 0-3.475-.32-4.721-.869-.623-.274-1.173-.62-1.579-1.04-.408-.426-.7-.965-.7-1.591v-9zM2.5 8V5.724c.241.15.503.286.779.407C4.525 6.68 6.195 7 8 7c1.805 0 3.475-.32 4.722-.869.275-.121.537-.257.778-.407V8c0 .133-.058.318-.282.55-.227.237-.592.484-1.1.708C11.101 9.705 9.646 10 8 10c-1.647 0-3.102-.295-4.117-.742-.51-.224-.874-.47-1.101-.707C2.558 8.318 2.5 8.133 2.5 8zm0 2.225V12.5c0 .133.058.318.282.55.227.237.592.484 1.1.708 1.016.447 2.471.742 4.118.742 1.647 0 3.102-.295 4.117-.742.51-.224.874-.47 1.101-.707.224-.233.282-.418.282-.551v-2.275c-.241.15-.503.285-.778.406-1.247.549-2.917.869-4.722.869-1.805 0-3.475-.32-4.721-.869a6.236 6.236 0 01-.779-.406z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-database%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 1.25c-2.487 0-4.774.402-6.466 1.079-.844.337-1.577.758-2.112 1.264C2.886 4.1 2.5 4.744 2.5 5.5v12.987l.026.013H2.5c0 .756.386 1.4.922 1.907.535.506 1.268.927 2.112 1.264 1.692.677 3.979 1.079 6.466 1.079s4.773-.402 6.466-1.079c.844-.337 1.577-.758 2.112-1.264.536-.507.922-1.151.922-1.907h-.026l.026-.013V5.5c0-.756-.386-1.4-.922-1.907-.535-.506-1.268-.927-2.112-1.264C16.773 1.652 14.487 1.25 12 1.25zM4 5.5c0-.21.104-.487.453-.817.35-.332.899-.666 1.638-.962C7.566 3.131 9.655 2.75 12 2.75c2.345 0 4.434.382 5.909.971.74.296 1.287.63 1.638.962.35.33.453.606.453.817 0 .21-.104.487-.453.817-.35.332-.899.666-1.638.962-1.475.59-3.564.971-5.909.971-2.345 0-4.434-.382-5.909-.971-.74-.296-1.287-.63-1.638-.962C4.103 5.987 4 5.711 4 5.5zM20 12V7.871a7.842 7.842 0 01-1.534.8C16.773 9.348 14.487 9.75 12 9.75s-4.774-.402-6.466-1.079A7.843 7.843 0 014 7.871V12c0 .21.104.487.453.817.35.332.899.666 1.638.961 1.475.59 3.564.972 5.909.972 2.345 0 4.434-.382 5.909-.972.74-.295 1.287-.629 1.638-.96.35-.33.453-.607.453-.818zM4 14.371c.443.305.963.572 1.534.8 1.692.677 3.979 1.079 6.466 1.079s4.773-.402 6.466-1.079a7.842 7.842 0 001.534-.8v4.116l.013.013H20c0 .21-.104.487-.453.817-.35.332-.899.666-1.638.962-1.475.59-3.564.971-5.909.971-2.345 0-4.434-.382-5.909-.971-.74-.296-1.287-.63-1.638-.962-.35-.33-.453-.606-.453-.817h-.013L4 18.487V14.37z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Dependabot returns a string representing an "dependabot" Octicon.
func Dependabot(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-dependabot%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M5.75 7.5a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5a.75.75 0 01.75-.75zm5.25.75a.75.75 0 00-1.5 0v1.5a.75.75 0 001.5 0v-1.5z\"></path><path fill-rule=\"evenodd\" d=\"M6.25 0a.75.75 0 000 1.5H7.5v2H3.75A2.25 2.25 0 001.5 5.75V8H.75a.75.75 0 000 1.5h.75v2.75a2.25 2.25 0 002.25 2.25h8.5a2.25 2.25 0 002.25-2.25V9.5h.75a.75.75 0 000-1.5h-.75V5.75a2.25 2.25 0 00-2.25-2.25H9V.75A.75.75 0 008.25 0h-2zM3 5.75A.75.75 0 013.75 5h8.5a.75.75 0 01.75.75v6.5a.75.75 0 01-.75.75h-8.5a.75.75 0 01-.75-.75v-6.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-dependabot%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M8.75 11a.75.75 0 01.75.75v3.5a.75.75 0 01-1.5 0v-3.5a.75.75 0 01.75-.75zm7.25.75a.75.75 0 00-1.5 0v3.5a.75.75 0 001.5 0v-3.5z\"></path><path fill-rule=\"evenodd\" d=\"M9.813 1a.75.75 0 000 1.5H11.5V5H4.25A2.25 2.25 0 002 7.25v5.25H.75a.75.75 0 000 1.5H2v5.75A2.25 2.25 0 004.25 22h15.5A2.25 2.25 0 0022 19.75V14h1.25a.75.75 0 000-1.5H22V7.25A2.25 2.25 0 0019.75 5H13V1.75a.75.75 0 00-.75-.75H9.812zM3.5 7.25a.75.75 0 01.75-.75h15.5a.75.75 0 01.75.75v12.5a.75.75 0 01-.75.75H4.25a.75.75 0 01-.75-.75V7.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DesktopDownload returns a string representing an "desktop-download" Octicon.
func DesktopDownload(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-desktop-download%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M4.927 5.427l2.896 2.896a.25.25 0 00.354 0l2.896-2.896A.25.25 0 0010.896 5H8.75V.75a.75.75 0 10-1.5 0V5H5.104a.25.25 0 00-.177.427z\"></path><path d=\"M1.573 2.573a.25.25 0 00-.073.177v7.5a.25.25 0 00.25.25h12.5a.25.25 0 00.25-.25v-7.5a.25.25 0 00-.25-.25h-3a.75.75 0 110-1.5h3A1.75 1.75 0 0116 2.75v7.5A1.75 1.75 0 0114.25 12h-3.727c.099 1.041.52 1.872 1.292 2.757A.75.75 0 0111.25 16h-6.5a.75.75 0 01-.565-1.243c.772-.885 1.192-1.716 1.292-2.757H1.75A1.75 1.75 0 010 10.25v-7.5A1.75 1.75 0 011.75 1h3a.75.75 0 010 1.5h-3a.25.25 0 00-.177.073zM6.982 12a5.72 5.72 0 01-.765 2.5h3.566a5.72 5.72 0 01-.765-2.5H6.982z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-desktop-download%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M11.25 9.331V.75a.75.75 0 011.5 0v8.58l1.949-2.11A.75.75 0 1115.8 8.237l-3.25 3.52a.75.75 0 01-1.102 0l-3.25-3.52A.75.75 0 119.3 7.22l1.949 2.111z\"></path><path fill-rule=\"evenodd\" d=\"M2.5 3.75a.25.25 0 01.25-.25h5.5a.75.75 0 100-1.5h-5.5A1.75 1.75 0 001 3.75v11.5c0 .966.784 1.75 1.75 1.75h6.204c-.171 1.375-.805 2.652-1.77 3.757A.75.75 0 007.75 22h8.5a.75.75 0 00.565-1.243c-.964-1.105-1.598-2.382-1.769-3.757h6.204A1.75 1.75 0 0023 15.25V3.75A1.75 1.75 0 0021.25 2h-5.5a.75.75 0 000 1.5h5.5a.25.25 0 01.25.25v11.5a.25.25 0 01-.25.25H2.75a.25.25 0 01-.25-.25V3.75zM10.463 17c-.126 1.266-.564 2.445-1.223 3.5h5.52c-.66-1.055-1.098-2.234-1.223-3.5h-3.074z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DeviceCamera returns a string representing an "device-camera" Octicon.
func DeviceCamera(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-device-camera%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M15 3H7c0-.55-.45-1-1-1H2c-.55 0-1 .45-1 1-.55 0-1 .45-1 1v9c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V4c0-.55-.45-1-1-1zM6 5H2V4h4v1zm4.5 7C8.56 12 7 10.44 7 8.5S8.56 5 10.5 5 14 6.56 14 8.5 12.44 12 10.5 12zM13 8.5c0 1.38-1.13 2.5-2.5 2.5S8 9.87 8 8.5 9.13 6 10.5 6 13 7.13 13 8.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DeviceCameraVideo returns a string representing an "device-camera-video" Octicon.
func DeviceCameraVideo(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-device-camera-video%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M16 3.75a.75.75 0 00-1.136-.643L11 5.425V4.75A1.75 1.75 0 009.25 3h-7.5A1.75 1.75 0 000 4.75v6.5C0 12.216.784 13 1.75 13h7.5A1.75 1.75 0 0011 11.25v-.675l3.864 2.318A.75.75 0 0016 12.25v-8.5zm-5 5.075l3.5 2.1v-5.85l-3.5 2.1v1.65zM9.5 6.75v-2a.25.25 0 00-.25-.25h-7.5a.25.25 0 00-.25.25v6.5c0 .138.112.25.25.25h7.5a.25.25 0 00.25-.25v-4.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-device-camera-video%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M24 5.25a.75.75 0 00-1.136-.643L16.5 8.425V6.25a1.75 1.75 0 00-1.75-1.75h-13A1.75 1.75 0 000 6.25v11C0 18.216.784 19 1.75 19h13a1.75 1.75 0 001.75-1.75v-2.175l6.364 3.818A.75.75 0 0024 18.25v-13zm-7.5 8.075l6 3.6V6.575l-6 3.6v3.15zM15 9.75v-3.5a.25.25 0 00-.25-.25h-13a.25.25 0 00-.25.25v11c0 .138.112.25.25.25h13a.25.25 0 00.25-.25v-7.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DeviceDesktop returns a string representing an "device-desktop" Octicon.
func DeviceDesktop(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-device-desktop%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 2.5h12.5a.25.25 0 01.25.25v7.5a.25.25 0 01-.25.25H1.75a.25.25 0 01-.25-.25v-7.5a.25.25 0 01.25-.25zM14.25 1H1.75A1.75 1.75 0 000 2.75v7.5C0 11.216.784 12 1.75 12h3.727c-.1 1.041-.52 1.872-1.292 2.757A.75.75 0 004.75 16h6.5a.75.75 0 00.565-1.243c-.772-.885-1.193-1.716-1.292-2.757h3.727A1.75 1.75 0 0016 10.25v-7.5A1.75 1.75 0 0014.25 1zM9.018 12H6.982a5.72 5.72 0 01-.765 2.5h3.566a5.72 5.72 0 01-.765-2.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-device-desktop%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.954 17H2.75A1.75 1.75 0 011 15.25V3.75C1 2.784 1.784 2 2.75 2h18.5c.966 0 1.75.784 1.75 1.75v11.5A1.75 1.75 0 0121.25 17h-6.204c.171 1.375.805 2.652 1.769 3.757A.75.75 0 0116.25 22h-8.5a.75.75 0 01-.565-1.243c.964-1.105 1.598-2.382 1.769-3.757zM21.5 3.75v11.5a.25.25 0 01-.25.25H2.75a.25.25 0 01-.25-.25V3.75a.25.25 0 01.25-.25h18.5a.25.25 0 01.25.25zM13.537 17c.125 1.266.564 2.445 1.223 3.5H9.24c.659-1.055 1.097-2.234 1.223-3.5h3.074z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DeviceMobile returns a string representing an "device-mobile" Octicon.
func DeviceMobile(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-device-mobile%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 0A1.75 1.75 0 002 1.75v12.5c0 .966.784 1.75 1.75 1.75h8.5A1.75 1.75 0 0014 14.25V1.75A1.75 1.75 0 0012.25 0h-8.5zM3.5 1.75a.25.25 0 01.25-.25h8.5a.25.25 0 01.25.25v12.5a.25.25 0 01-.25.25h-8.5a.25.25 0 01-.25-.25V1.75zM8 13a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-device-mobile%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M10.25 5.25A.75.75 0 0111 4.5h2A.75.75 0 0113 6h-2a.75.75 0 01-.75-.75zM12 19.5a1 1 0 100-2 1 1 0 000 2z\"></path><path fill-rule=\"evenodd\" d=\"M4 2.75C4 1.784 4.784 1 5.75 1h12.5c.966 0 1.75.784 1.75 1.75v18.5A1.75 1.75 0 0118.25 23H5.75A1.75 1.75 0 014 21.25V2.75zm1.75-.25a.25.25 0 00-.25.25v18.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V2.75a.25.25 0 00-.25-.25H5.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Diamond returns a string representing an "diamond" Octicon.
func Diamond(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-diamond%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M.527 9.237a1.75 1.75 0 010-2.474L6.777.512a1.75 1.75 0 012.475 0l6.251 6.25a1.751 1.751 0 010 2.475l-6.25 6.251a1.751 1.751 0 01-2.475 0L.527 9.238v-.001zm1.06-1.414a.25.25 0 000 .354l6.251 6.25a.25.25 0 00.354 0l6.25-6.25a.25.25 0 000-.354l-6.25-6.25a.25.25 0 00-.354 0l-6.25 6.25h-.001z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-diamond%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.527 13.237a1.75 1.75 0 010-2.474l9.272-9.273a1.75 1.75 0 012.475 0l9.272 9.273a1.75 1.75 0 010 2.474l-9.272 9.272a1.75 1.75 0 01-2.475 0l-9.272-9.272zm1.06-1.414a.25.25 0 000 .354l9.273 9.272a.25.25 0 00.353 0l9.272-9.272a.25.25 0 000-.354l-9.272-9.272a.25.25 0 00-.353 0l-9.273 9.272z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Diff returns a string representing an "diff" Octicon.
func Diff(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-diff%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.75 1.75a.75.75 0 00-1.5 0V5H4a.75.75 0 000 1.5h3.25v3.25a.75.75 0 001.5 0V6.5H12A.75.75 0 0012 5H8.75V1.75zM4 13a.75.75 0 000 1.5h8a.75.75 0 100-1.5H4z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-diff%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.25 3.5a.75.75 0 01.75.75V8.5h4.25a.75.75 0 010 1.5H13v4.25a.75.75 0 01-1.5 0V10H7.25a.75.75 0 010-1.5h4.25V4.25a.75.75 0 01.75-.75zM6.562 19.25a.75.75 0 01.75-.75h9.938a.75.75 0 010 1.5H7.312a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DiffAdded returns a string representing an "diff-added" Octicon.
func DiffAdded(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-diff-added%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13.25 2.5H2.75a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h10.5a.25.25 0 00.25-.25V2.75a.25.25 0 00-.25-.25zM2.75 1h10.5c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0113.25 15H2.75A1.75 1.75 0 011 13.25V2.75C1 1.784 1.784 1 2.75 1zM8 4a.75.75 0 01.75.75v2.5h2.5a.75.75 0 010 1.5h-2.5v2.5a.75.75 0 01-1.5 0v-2.5h-2.5a.75.75 0 010-1.5h2.5v-2.5A.75.75 0 018 4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DiffIgnored returns a string representing an "diff-ignored" Octicon.
func DiffIgnored(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-diff-ignored%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 2.5h10.5a.25.25 0 01.25.25v10.5a.25.25 0 01-.25.25H2.75a.25.25 0 01-.25-.25V2.75a.25.25 0 01.25-.25zM13.25 1H2.75A1.75 1.75 0 001 2.75v10.5c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0015 13.25V2.75A1.75 1.75 0 0013.25 1zm-1.97 4.78a.75.75 0 00-1.06-1.06l-5.5 5.5a.75.75 0 101.06 1.06l5.5-5.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DiffModified returns a string representing an "diff-modified" Octicon.
func DiffModified(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-diff-modified%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 2.5h10.5a.25.25 0 01.25.25v10.5a.25.25 0 01-.25.25H2.75a.25.25 0 01-.25-.25V2.75a.25.25 0 01.25-.25zM13.25 1H2.75A1.75 1.75 0 001 2.75v10.5c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0015 13.25V2.75A1.75 1.75 0 0013.25 1zM8 10a2 2 0 100-4 2 2 0 000 4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DiffRemoved returns a string representing an "diff-removed" Octicon.
func DiffRemoved(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-diff-removed%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 2.5h10.5a.25.25 0 01.25.25v10.5a.25.25 0 01-.25.25H2.75a.25.25 0 01-.25-.25V2.75a.25.25 0 01.25-.25zM13.25 1H2.75A1.75 1.75 0 001 2.75v10.5c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0015 13.25V2.75A1.75 1.75 0 0013.25 1zm-2 7.75a.75.75 0 000-1.5h-6.5a.75.75 0 000 1.5h6.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DiffRenamed returns a string representing an "diff-renamed" Octicon.
func DiffRenamed(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-diff-renamed%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 2.5h10.5a.25.25 0 01.25.25v10.5a.25.25 0 01-.25.25H2.75a.25.25 0 01-.25-.25V2.75a.25.25 0 01.25-.25zM13.25 1H2.75A1.75 1.75 0 001 2.75v10.5c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0015 13.25V2.75A1.75 1.75 0 0013.25 1zm-1.47 7.53a.75.75 0 000-1.06L8.53 4.22a.75.75 0 00-1.06 1.06l1.97 1.97H4.75a.75.75 0 000 1.5h4.69l-1.97 1.97a.75.75 0 101.06 1.06l3.25-3.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Dot returns a string representing an "dot" Octicon.
func Dot(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-dot%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 5.5a2.5 2.5 0 100 5 2.5 2.5 0 000-5zM4 8a4 4 0 118 0 4 4 0 01-8 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-dot%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 16.5a4.5 4.5 0 100-9 4.5 4.5 0 000 9zm0 1.5a6 6 0 100-12 6 6 0 000 12z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// DotFill returns a string representing an "dot-fill" Octicon.
func DotFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-dot-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 4a4 4 0 100 8 4 4 0 000-8z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-dot-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12 18a6 6 0 100-12 6 6 0 000 12z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Download returns a string representing an "download" Octicon.
func Download(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-download%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.47 10.78a.75.75 0 001.06 0l3.75-3.75a.75.75 0 00-1.06-1.06L8.75 8.44V1.75a.75.75 0 00-1.5 0v6.69L4.78 5.97a.75.75 0 00-1.06 1.06l3.75 3.75zM3.75 13a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-download%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M4.97 11.03a.75.75 0 111.06-1.06L11 14.94V2.75a.75.75 0 011.5 0v12.19l4.97-4.97a.75.75 0 111.06 1.06l-6.25 6.25a.75.75 0 01-1.06 0l-6.25-6.25zm-.22 9.47a.75.75 0 000 1.5h14.5a.75.75 0 000-1.5H4.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Duplicate returns a string representing an "duplicate" Octicon.
func Duplicate(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-duplicate%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M10.5 3a.75.75 0 01.75.75v1h1a.75.75 0 010 1.5h-1v1a.75.75 0 01-1.5 0v-1h-1a.75.75 0 010-1.5h1v-1A.75.75 0 0110.5 3z\"></path><path fill-rule=\"evenodd\" d=\"M6.75 0A1.75 1.75 0 005 1.75v7.5c0 .966.784 1.75 1.75 1.75h7.5A1.75 1.75 0 0016 9.25v-7.5A1.75 1.75 0 0014.25 0h-7.5zM6.5 1.75a.25.25 0 01.25-.25h7.5a.25.25 0 01.25.25v7.5a.25.25 0 01-.25.25h-7.5a.25.25 0 01-.25-.25v-7.5z\"></path><path d=\"M1.75 5A1.75 1.75 0 000 6.75v7.5C0 15.216.784 16 1.75 16h7.5A1.75 1.75 0 0011 14.25v-1.5a.75.75 0 00-1.5 0v1.5a.25.25 0 01-.25.25h-7.5a.25.25 0 01-.25-.25v-7.5a.25.25 0 01.25-.25h1.5a.75.75 0 000-1.5h-1.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-duplicate%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M14.513 6a.75.75 0 01.75.75v2h1.987a.75.75 0 010 1.5h-1.987v2a.75.75 0 11-1.5 0v-2H11.75a.75.75 0 010-1.5h2.013v-2a.75.75 0 01.75-.75z\"></path><path fill-rule=\"evenodd\" d=\"M7.024 3.75c0-.966.784-1.75 1.75-1.75H20.25c.966 0 1.75.784 1.75 1.75v11.498a1.75 1.75 0 01-1.75 1.75H8.774a1.75 1.75 0 01-1.75-1.75V3.75zm1.75-.25a.25.25 0 00-.25.25v11.498c0 .139.112.25.25.25H20.25a.25.25 0 00.25-.25V3.75a.25.25 0 00-.25-.25H8.774z\"></path><path d=\"M1.995 10.749a1.75 1.75 0 011.75-1.751H5.25a.75.75 0 110 1.5H3.745a.25.25 0 00-.25.25L3.5 20.25c0 .138.111.25.25.25h9.5a.25.25 0 00.25-.25v-1.51a.75.75 0 111.5 0v1.51A1.75 1.75 0 0113.25 22h-9.5A1.75 1.75 0 012 20.25l-.005-9.501z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Ellipsis returns a string representing an "ellipsis" Octicon.
func Ellipsis(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-ellipsis%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 5.75C0 4.784.784 4 1.75 4h12.5c.966 0 1.75.784 1.75 1.75v4.5A1.75 1.75 0 0114.25 12H1.75A1.75 1.75 0 010 10.25v-4.5zM4 7a1 1 0 100 2 1 1 0 000-2zm3 1a1 1 0 112 0 1 1 0 01-2 0zm5-1a1 1 0 100 2 1 1 0 000-2z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Eye returns a string representing an "eye" Octicon.
func Eye(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-eye%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.679 7.932c.412-.621 1.242-1.75 2.366-2.717C5.175 4.242 6.527 3.5 8 3.5c1.473 0 2.824.742 3.955 1.715 1.124.967 1.954 2.096 2.366 2.717a.119.119 0 010 .136c-.412.621-1.242 1.75-2.366 2.717C10.825 11.758 9.473 12.5 8 12.5c-1.473 0-2.824-.742-3.955-1.715C2.92 9.818 2.09 8.69 1.679 8.068a.119.119 0 010-.136zM8 2c-1.981 0-3.67.992-4.933 2.078C1.797 5.169.88 6.423.43 7.1a1.619 1.619 0 000 1.798c.45.678 1.367 1.932 2.637 3.024C4.329 13.008 6.019 14 8 14c1.981 0 3.67-.992 4.933-2.078 1.27-1.091 2.187-2.345 2.637-3.023a1.619 1.619 0 000-1.798c-.45-.678-1.367-1.932-2.637-3.023C11.671 2.992 9.981 2 8 2zm0 8a2 2 0 100-4 2 2 0 000 4z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-eye%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M15.5 12a3.5 3.5 0 11-7 0 3.5 3.5 0 017 0z\"></path><path fill-rule=\"evenodd\" d=\"M12 3.5c-3.432 0-6.125 1.534-8.054 3.24C2.02 8.445.814 10.352.33 11.202a1.6 1.6 0 000 1.598c.484.85 1.69 2.758 3.616 4.46C5.876 18.966 8.568 20.5 12 20.5c3.432 0 6.125-1.534 8.054-3.24 1.926-1.704 3.132-3.611 3.616-4.461a1.6 1.6 0 000-1.598c-.484-.85-1.69-2.757-3.616-4.46C18.124 5.034 15.432 3.5 12 3.5zM1.633 11.945c.441-.774 1.551-2.528 3.307-4.08C6.69 6.314 9.045 5 12 5c2.955 0 5.309 1.315 7.06 2.864 1.756 1.553 2.866 3.307 3.307 4.08a.111.111 0 01.017.056.111.111 0 01-.017.056c-.441.774-1.551 2.527-3.307 4.08C17.31 17.685 14.955 19 12 19c-2.955 0-5.309-1.315-7.06-2.864-1.756-1.553-2.866-3.306-3.307-4.08A.11.11 0 011.616 12a.11.11 0 01.017-.055z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// EyeClosed returns a string representing an "eye-closed" Octicon.
func EyeClosed(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-eye-closed%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M.143 2.31a.75.75 0 011.047-.167l14.5 10.5a.75.75 0 11-.88 1.214l-2.248-1.628C11.346 13.19 9.792 14 8 14c-1.981 0-3.67-.992-4.933-2.078C1.797 10.832.88 9.577.43 8.9a1.618 1.618 0 010-1.797c.353-.533.995-1.42 1.868-2.305L.31 3.357A.75.75 0 01.143 2.31zm3.386 3.378a14.21 14.21 0 00-1.85 2.244.12.12 0 00-.022.068c0 .021.006.045.022.068.412.621 1.242 1.75 2.366 2.717C5.175 11.758 6.527 12.5 8 12.5c1.195 0 2.31-.488 3.29-1.191L9.063 9.695A2 2 0 016.058 7.52l-2.53-1.832zM8 3.5c-.516 0-1.017.09-1.499.251a.75.75 0 11-.473-1.423A6.23 6.23 0 018 2c1.981 0 3.67.992 4.933 2.078 1.27 1.091 2.187 2.345 2.637 3.023a1.619 1.619 0 010 1.798c-.11.166-.248.365-.41.587a.75.75 0 11-1.21-.887c.148-.201.272-.382.371-.53a.119.119 0 000-.137c-.412-.621-1.242-1.75-2.366-2.717C10.825 4.242 9.473 3.5 8 3.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-eye-closed%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M8.052 5.837A9.715 9.715 0 0112 5c2.955 0 5.309 1.315 7.06 2.864 1.756 1.553 2.866 3.307 3.307 4.08a.11.11 0 01.016.055.122.122 0 01-.017.06 16.766 16.766 0 01-1.53 2.218.75.75 0 101.163.946 18.253 18.253 0 001.67-2.42 1.607 1.607 0 00.001-1.602c-.485-.85-1.69-2.757-3.616-4.46C18.124 5.034 15.432 3.5 12 3.5c-1.695 0-3.215.374-4.552.963a.75.75 0 00.604 1.373z\"></path><path fill-rule=\"evenodd\" d=\"M19.166 17.987C17.328 19.38 14.933 20.5 12 20.5c-3.432 0-6.125-1.534-8.054-3.24C2.02 15.556.814 13.648.33 12.798a1.606 1.606 0 01.001-1.6A18.305 18.305 0 013.648 7.01L1.317 5.362a.75.75 0 11.866-1.224l20.5 14.5a.75.75 0 11-.866 1.224l-2.651-1.875zM4.902 7.898c-1.73 1.541-2.828 3.273-3.268 4.044a.118.118 0 00-.017.059c0 .015.003.034.016.055.441.774 1.551 2.527 3.307 4.08C6.69 17.685 9.045 19 12 19c2.334 0 4.29-.82 5.874-1.927l-3.516-2.487a3.5 3.5 0 01-5.583-3.949L4.902 7.899z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedDiscussion returns a string representing an "feed-discussion" Octicon.
func FeedDiscussion(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-discussion%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zM4 5a1 1 0 011-1h6a1 1 0 011 1v5a1 1 0 01-1 1H8.707l-1.853 1.854A.5.5 0 016 12.5V11H5a1 1 0 01-1-1V5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedForked returns a string representing an "feed-forked" Octicon.
func FeedForked(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-forked%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zM6 6.928a1.75 1.75 0 10-1 0V7.5A1.5 1.5 0 006.5 9h1v1.072a1.75 1.75 0 101 0V9h1A1.5 1.5 0 0011 7.5v-.572a1.75 1.75 0 10-1 0V7.5a.5.5 0 01-.5.5h-3a.5.5 0 01-.5-.5v-.572z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedHeart returns a string representing an "feed-heart" Octicon.
func FeedHeart(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-heart%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zm2.33-11.5c-1.22 0-1.83.5-2.323 1.136C7.513 5 6.903 4.5 5.682 4.5c-1.028 0-2.169.784-2.169 2.5 0 1.499 1.493 3.433 3.246 4.517.52.321.89.479 1.248.484.357-.005.728-.163 1.247-.484C11.007 10.433 12.5 8.5 12.5 7c0-1.716-1.14-2.5-2.17-2.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedMerged returns a string representing an "feed-merged" Octicon.
func FeedMerged(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-merged%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zm.25-11.25a1.75 1.75 0 01-1.207 1.664A2 2 0 009 8h.571a1.75 1.75 0 110 1H9a2.99 2.99 0 01-2-.764v1.336a1.75 1.75 0 11-1 0V6.428A1.75 1.75 0 118.25 4.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedPerson returns a string representing an "feed-person" Octicon.
func FeedPerson(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-person%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zm.847-8.145a2.502 2.502 0 10-1.694 0C5.471 8.261 4 9.775 4 11c0 .395.145.995 1 .995h6c.855 0 1-.6 1-.995 0-1.224-1.47-2.74-3.153-3.145z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedRepo returns a string representing an "feed-repo" Octicon.
func FeedRepo(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-repo%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zM5.5 4A1.5 1.5 0 004 5.5v5c0 .828.5 1.5 1 1.5v-1a1 1 0 011-1h5v1h-1v1h1.5a.5.5 0 00.5-.5v-7a.5.5 0 00-.5-.5h-6zm.5 7.25a.25.25 0 01.25-.25H9v2.764a.25.25 0 01-.426.178l-.898-.888a.25.25 0 00-.352 0l-.898.888A.25.25 0 016 13.764V11.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedRocket returns a string representing an "feed-rocket" Octicon.
func FeedRocket(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-rocket%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zm3.031-12a4.38 4.38 0 00-3.097 1.283l-.23.229c-.156.157-.308.32-.452.49H5.65a.876.876 0 00-.746.417l-.856 1.388a.375.375 0 00.21.556l1.552.477 1.35 1.35.478 1.553a.375.375 0 00.555.21l1.389-.855a.876.876 0 00.416-.746V8.747c.17-.144.333-.295.49-.452l.23-.23A4.38 4.38 0 0012 4.969v-.093A.876.876 0 0011.124 4h-.093zm-5.107 7.144a.81.81 0 01-.188.263c-.394.394-1.258.563-1.62.619a.124.124 0 01-.143-.143c.056-.362.225-1.226.62-1.62a.808.808 0 011.33.881z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedStar returns a string representing an "feed-star" Octicon.
func FeedStar(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-star%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zm.252-12.932a.478.478 0 00-.682.195l-1.2 2.432-2.684.39a.478.478 0 00-.266.816l1.944 1.892-.46 2.674a.478.478 0 00.694.504L8 10.709l2.4 1.261a.478.478 0 00.694-.504l-.458-2.673L12.578 6.9a.479.479 0 00-.265-.815l-2.685-.39-1.2-2.432a.478.478 0 00-.176-.195z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedTag returns a string representing an "feed-tag" Octicon.
func FeedTag(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-tag%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M7.22 6.5a.72.72 0 11-1.44 0 .72.72 0 011.44 0z\"></path><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zM4 8.379V5a1 1 0 011-1h3.379a1.5 1.5 0 011.06.44l3.213 3.211a1.2 1.2 0 010 1.698l-3.303 3.303a1.2 1.2 0 01-1.698 0L4.44 9.439A1.5 1.5 0 014 8.38z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FeedTrophy returns a string representing an "feed-trophy" Octicon.
func FeedTrophy(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-feed-trophy%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M11 5h1v1.146a1 1 0 01-.629.928L11 7.223V5zM5 7.223l-.371-.149A1 1 0 014 6.146V5h1v2.223z\"></path><path fill-rule=\"evenodd\" d=\"M8 16A8 8 0 108 0a8 8 0 000 16zM3 5a1 1 0 011-1h8a1 1 0 011 1v1.146a2 2 0 01-1.257 1.857l-.865.346a3.005 3.005 0 01-2.294 2.094C8.78 11.405 9.342 12 10.5 12a.5.5 0 010 1h-5a.5.5 0 010-1h.002c1.156 0 1.718-.596 1.914-1.557A3.005 3.005 0 015.122 8.35l-.865-.346A2 2 0 013 6.146V5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// File returns a string representing an "file" Octicon.
func File(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 1.5a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h9.5a.25.25 0 00.25-.25V6h-2.75A1.75 1.75 0 019 4.25V1.5H3.75zm6.75.062V4.25c0 .138.112.25.25.25h2.688a.252.252 0 00-.011-.013l-2.914-2.914a.272.272 0 00-.013-.011zM2 1.75C2 .784 2.784 0 3.75 0h6.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v9.586A1.75 1.75 0 0113.25 16h-9.5A1.75 1.75 0 012 14.25V1.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5 2.5a.5.5 0 00-.5.5v18a.5.5 0 00.5.5h14a.5.5 0 00.5-.5V8.5h-4a2 2 0 01-2-2v-4H5zm10 0v4a.5.5 0 00.5.5h4a.5.5 0 00-.146-.336l-4.018-4.018A.5.5 0 0015 2.5zM3 3a2 2 0 012-2h9.982a2 2 0 011.414.586l4.018 4.018A2 2 0 0121 7.018V21a2 2 0 01-2 2H5a2 2 0 01-2-2V3z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileAdded returns a string representing an "file-added" Octicon.
func FileAdded(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-added%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 1.5a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h9.5a.25.25 0 00.25-.25V4.664a.25.25 0 00-.073-.177l-2.914-2.914a.25.25 0 00-.177-.073H3.75zM2 1.75C2 .784 2.784 0 3.75 0h6.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v9.586A1.75 1.75 0 0113.25 16h-9.5A1.75 1.75 0 012 14.25V1.75zm6.23 3.508a.75.75 0 01.755.745l.01 1.497h1.497a.75.75 0 010 1.5H9v1.507a.75.75 0 01-1.5 0V9.005l-1.502.01a.75.75 0 11-.01-1.5l1.507-.01-.01-1.492a.75.75 0 01.745-.755z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileBadge returns a string representing an "file-badge" Octicon.
func FileBadge(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-badge%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M2.75 1.5a.25.25 0 00-.25.25v11.5c0 .138.112.25.25.25h3.5a.75.75 0 010 1.5h-3.5A1.75 1.75 0 011 13.25V1.75C1 .784 1.784 0 2.75 0h8a1.75 1.75 0 011.508.862.75.75 0 11-1.289.768.25.25 0 00-.219-.13h-8z\"></path><path fill-rule=\"evenodd\" d=\"M8 7a4 4 0 116.49 3.13l.995 4.973a.75.75 0 01-.991.852l-2.409-.876a.25.25 0 00-.17 0l-2.409.876a.75.75 0 01-.991-.852l.994-4.973A3.993 3.993 0 018 7zm4-2.5a2.5 2.5 0 100 5 2.5 2.5 0 000-5zm0 6.5a4 4 0 001.104-.154l.649 3.243-1.155-.42c-.386-.14-.81-.14-1.196 0l-1.155.42.649-3.243A4 4 0 0012 11z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileBinary returns a string representing an "file-binary" Octicon.
func FileBinary(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-binary%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4 1.75C4 .784 4.784 0 5.75 0h5.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v8.586A1.75 1.75 0 0114.25 15h-9a.75.75 0 010-1.5h9a.25.25 0 00.25-.25V6h-2.75A1.75 1.75 0 0110 4.25V1.5H5.75a.25.25 0 00-.25.25v2a.75.75 0 01-1.5 0v-2zm7.5-.188V4.25c0 .138.112.25.25.25h2.688a.252.252 0 00-.011-.013l-2.914-2.914a.272.272 0 00-.013-.011zM0 7.75C0 6.784.784 6 1.75 6h1.5C4.216 6 5 6.784 5 7.75v2.5A1.75 1.75 0 013.25 12h-1.5A1.75 1.75 0 010 10.25v-2.5zm1.75-.25a.25.25 0 00-.25.25v2.5c0 .138.112.25.25.25h1.5a.25.25 0 00.25-.25v-2.5a.25.25 0 00-.25-.25h-1.5zm5-1.5a.75.75 0 000 1.5h.75v3h-.75a.75.75 0 000 1.5h3a.75.75 0 000-1.5H9V6.75A.75.75 0 008.25 6h-1.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-binary%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3 3a2 2 0 012-2h9.982a2 2 0 011.414.586l4.018 4.018A2 2 0 0121 7.018V21a2 2 0 01-2 2H4.75a.75.75 0 010-1.5H19a.5.5 0 00.5-.5V8.5h-4a2 2 0 01-2-2v-4H5a.5.5 0 00-.5.5v6.25a.75.75 0 01-1.5 0V3zm12-.5v4a.5.5 0 00.5.5h4a.5.5 0 00-.146-.336l-4.018-4.018A.5.5 0 0015 2.5z\"></path><path fill-rule=\"evenodd\" d=\"M0 13.75C0 12.784.784 12 1.75 12h3c.966 0 1.75.784 1.75 1.75v4a1.75 1.75 0 01-1.75 1.75h-3A1.75 1.75 0 010 17.75v-4zm1.75-.25a.25.25 0 00-.25.25v4c0 .138.112.25.25.25h3a.25.25 0 00.25-.25v-4a.25.25 0 00-.25-.25h-3z\"></path><path d=\"M9 12a.75.75 0 000 1.5h1.5V18H9a.75.75 0 000 1.5h4.5a.75.75 0 000-1.5H12v-5.25a.75.75 0 00-.75-.75H9z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileCode returns a string representing an "file-code" Octicon.
func FileCode(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-code%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4 1.75C4 .784 4.784 0 5.75 0h5.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v8.586A1.75 1.75 0 0114.25 15h-9a.75.75 0 010-1.5h9a.25.25 0 00.25-.25V6h-2.75A1.75 1.75 0 0110 4.25V1.5H5.75a.25.25 0 00-.25.25v2.5a.75.75 0 01-1.5 0v-2.5zm7.5-.188V4.25c0 .138.112.25.25.25h2.688a.252.252 0 00-.011-.013l-2.914-2.914a.272.272 0 00-.013-.011zM5.72 6.72a.75.75 0 000 1.06l1.47 1.47-1.47 1.47a.75.75 0 101.06 1.06l2-2a.75.75 0 000-1.06l-2-2a.75.75 0 00-1.06 0zM3.28 7.78a.75.75 0 00-1.06-1.06l-2 2a.75.75 0 000 1.06l2 2a.75.75 0 001.06-1.06L1.81 9.25l1.47-1.47z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-code%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3 3a2 2 0 012-2h9.982a2 2 0 011.414.586l4.018 4.018A2 2 0 0121 7.018V21a2 2 0 01-2 2H4.75a.75.75 0 010-1.5H19a.5.5 0 00.5-.5V8.5h-4a2 2 0 01-2-2v-4H5a.5.5 0 00-.5.5v6.25a.75.75 0 01-1.5 0V3zm12-.5v4a.5.5 0 00.5.5h4a.5.5 0 00-.146-.336l-4.018-4.018A.5.5 0 0015 2.5z\"></path><path d=\"M4.53 12.24a.75.75 0 01-.039 1.06l-2.639 2.45 2.64 2.45a.75.75 0 11-1.022 1.1l-3.23-3a.75.75 0 010-1.1l3.23-3a.75.75 0 011.06.04zm3.979 1.06a.75.75 0 111.02-1.1l3.231 3a.75.75 0 010 1.1l-3.23 3a.75.75 0 11-1.021-1.1l2.639-2.45-2.64-2.45z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileDiff returns a string representing an "file-diff" Octicon.
func FileDiff(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-diff%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 1.5a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h10.5a.25.25 0 00.25-.25V4.664a.25.25 0 00-.073-.177l-2.914-2.914a.25.25 0 00-.177-.073H2.75zM1 1.75C1 .784 1.784 0 2.75 0h7.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v9.586A1.75 1.75 0 0113.25 16H2.75A1.75 1.75 0 011 14.25V1.75zm7 1.5a.75.75 0 01.75.75v1.5h1.5a.75.75 0 010 1.5h-1.5v1.5a.75.75 0 01-1.5 0V7h-1.5a.75.75 0 010-1.5h1.5V4A.75.75 0 018 3.25zm-3 8a.75.75 0 01.75-.75h4.5a.75.75 0 010 1.5h-4.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-diff%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12.5 6.75a.75.75 0 00-1.5 0V9H8.75a.75.75 0 000 1.5H11v2.25a.75.75 0 001.5 0V10.5h2.25a.75.75 0 000-1.5H12.5V6.75zM8.75 16a.75.75 0 000 1.5h6a.75.75 0 000-1.5h-6z\"></path><path fill-rule=\"evenodd\" d=\"M5 1a2 2 0 00-2 2v18a2 2 0 002 2h14a2 2 0 002-2V7.018a2 2 0 00-.586-1.414l-4.018-4.018A2 2 0 0014.982 1H5zm-.5 2a.5.5 0 01.5-.5h9.982a.5.5 0 01.354.146l4.018 4.018a.5.5 0 01.146.354V21a.5.5 0 01-.5.5H5a.5.5 0 01-.5-.5V3z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileDirectory returns a string representing an "file-directory" Octicon.
func FileDirectory(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-directory%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 2.5a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25v-8.5a.25.25 0 00-.25-.25H7.5c-.55 0-1.07-.26-1.4-.7l-.9-1.2a.25.25 0 00-.2-.1H1.75zM0 2.75C0 1.784.784 1 1.75 1H5c.55 0 1.07.26 1.4.7l.9 1.2a.25.25 0 00.2.1h6.75c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0114.25 15H1.75A1.75 1.75 0 010 13.25V2.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-directory%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 4.5a.25.25 0 00-.25.25v14.5c0 .138.112.25.25.25h16.5a.25.25 0 00.25-.25V7.687a.25.25 0 00-.25-.25h-8.471a1.75 1.75 0 01-1.447-.765L8.928 4.61a.25.25 0 00-.208-.11H3.75zM2 4.75C2 3.784 2.784 3 3.75 3h4.971c.58 0 1.12.286 1.447.765l1.404 2.063a.25.25 0 00.207.11h8.471c.966 0 1.75.783 1.75 1.75V19.25A1.75 1.75 0 0120.25 21H3.75A1.75 1.75 0 012 19.25V4.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileDirectoryFill returns a string representing an "file-directory-fill" Octicon.
func FileDirectoryFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-directory-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M1.75 1A1.75 1.75 0 000 2.75v10.5C0 14.216.784 15 1.75 15h12.5A1.75 1.75 0 0016 13.25v-8.5A1.75 1.75 0 0014.25 3H7.5a.25.25 0 01-.2-.1l-.9-1.2C6.07 1.26 5.55 1 5 1H1.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-directory-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 4.75C2 3.784 2.784 3 3.75 3h4.971c.58 0 1.12.286 1.447.765l1.404 2.063a.25.25 0 00.207.11h8.471c.966 0 1.75.783 1.75 1.75V19.25A1.75 1.75 0 0120.25 21H3.75A1.75 1.75 0 012 19.25V4.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileDirectoryOpenFill returns a string representing an "file-directory-open-fill" Octicon.
func FileDirectoryOpenFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-directory-open-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M.513 1.513A1.75 1.75 0 011.75 1h3.5c.55 0 1.07.26 1.4.7l.9 1.2a.25.25 0 00.2.1H13a1 1 0 011 1v.5H2.75a.75.75 0 000 1.5h11.978a1 1 0 01.994 1.117L15 13.25A1.75 1.75 0 0113.25 15H1.75A1.75 1.75 0 010 13.25V2.75c0-.464.184-.91.513-1.237z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileMedia returns a string representing an "file-media" Octicon.
func FileMedia(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-media%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.25 4a.25.25 0 00-.25.25v15.5c0 .138.112.25.25.25h3.178L14 10.977a1.75 1.75 0 012.506-.032L22 16.44V4.25a.25.25 0 00-.25-.25H2.25zm3.496 17.5H21.75a1.75 1.75 0 001.75-1.75V4.25a1.75 1.75 0 00-1.75-1.75H2.25A1.75 1.75 0 00.5 4.25v15.5c0 .966.784 1.75 1.75 1.75h3.496zM22 19.75v-1.19l-6.555-6.554a.25.25 0 00-.358.004L7.497 20H21.75a.25.25 0 00.25-.25zM9 9.25a1.75 1.75 0 11-3.5 0 1.75 1.75 0 013.5 0zm1.5 0a3.25 3.25 0 11-6.5 0 3.25 3.25 0 016.5 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileMoved returns a string representing an "file-moved" Octicon.
func FileMoved(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-moved%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M2 1.75C2 .784 2.784 0 3.75 0h6.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v9.586A1.75 1.75 0 0113.25 16h-3.5a.75.75 0 010-1.5h3.5a.25.25 0 00.25-.25V4.664a.25.25 0 00-.073-.177l-2.914-2.914a.25.25 0 00-.177-.073H3.75a.25.25 0 00-.25.25v6.5a.75.75 0 01-1.5 0v-6.5z\"></path><path d=\"M5.427 15.573l3.146-3.146a.25.25 0 000-.354L5.427 8.927A.25.25 0 005 9.104V11.5H.75a.75.75 0 000 1.5H5v2.396c0 .223.27.335.427.177z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileRemoved returns a string representing an "file-removed" Octicon.
func FileRemoved(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-removed%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 1.5a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h9.5a.25.25 0 00.25-.25V4.664a.25.25 0 00-.073-.177l-2.914-2.914a.25.25 0 00-.177-.073H3.75zM2 1.75C2 .784 2.784 0 3.75 0h6.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v9.586A1.75 1.75 0 0113.25 16h-9.5A1.75 1.75 0 012 14.25V1.75zM8.25 7.5h2.242a.75.75 0 010 1.5h-2.24l-2.254.015a.75.75 0 01-.01-1.5L8.25 7.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileSubmodule returns a string representing an "file-submodule" Octicon.
func FileSubmodule(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-submodule%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 2.75C0 1.784.784 1 1.75 1H5c.55 0 1.07.26 1.4.7l.9 1.2a.25.25 0 00.2.1h6.75c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0114.25 15H1.75A1.75 1.75 0 010 13.25V2.75zm9.42 9.36l2.883-2.677a.25.25 0 000-.366L9.42 6.39a.25.25 0 00-.42.183V8.5H4.75a.75.75 0 100 1.5H9v1.927c0 .218.26.331.42.183z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-submodule%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 4.75C2 3.784 2.784 3 3.75 3h4.965a1.75 1.75 0 011.456.78l1.406 2.109a.25.25 0 00.208.111h8.465c.966 0 1.75.784 1.75 1.75v11.5A1.75 1.75 0 0120.25 21H3.75A1.75 1.75 0 012 19.25V4.75zm12.78 4.97a.75.75 0 10-1.06 1.06l1.72 1.72H6.75a.75.75 0 000 1.5h8.69l-1.72 1.72a.75.75 0 101.06 1.06l3-3a.75.75 0 000-1.06l-3-3z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileSymlinkFile returns a string representing an "file-symlink-file" Octicon.
func FileSymlinkFile(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-symlink-file%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 1.75C2 .784 2.784 0 3.75 0h5.586c.464 0 .909.184 1.237.513l2.914 2.914c.329.328.513.773.513 1.237v8.586A1.75 1.75 0 0112.25 15h-7a.75.75 0 010-1.5h7a.25.25 0 00.25-.25V6H9.75A1.75 1.75 0 018 4.25V1.5H3.75a.25.25 0 00-.25.25V4.5a.75.75 0 01-1.5 0V1.75zm7.5-.188V4.25c0 .138.112.25.25.25h2.688a.252.252 0 00-.011-.013L9.513 1.573a.248.248 0 00-.013-.011zm-8 10.675a2.25 2.25 0 012.262-2.25L4 9.99v1.938c0 .218.26.331.42.183l2.883-2.677a.25.25 0 000-.366L4.42 6.39a.25.25 0 00-.42.183V8.49l-.23-.001A3.75 3.75 0 000 12.238v1.012a.75.75 0 001.5 0v-1.013z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-symlink-file%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3 3a2 2 0 012-2h9.982a2 2 0 011.414.586l4.018 4.018A2 2 0 0121 7.018V21a2 2 0 01-2 2H4.75a.75.75 0 010-1.5H19a.5.5 0 00.5-.5V8.5h-4a2 2 0 01-2-2v-4H5a.5.5 0 00-.5.5v6.25a.75.75 0 01-1.5 0V3zm12-.5v4a.5.5 0 00.5.5h4a.5.5 0 00-.146-.336l-4.018-4.018A.5.5 0 0015 2.5zm-5.692 12l-2.104-2.236a.75.75 0 111.092-1.028l3.294 3.5a.75.75 0 010 1.028l-3.294 3.5a.75.75 0 11-1.092-1.028L9.308 16H4.09a2.59 2.59 0 00-2.59 2.59v3.16a.75.75 0 01-1.5 0v-3.16a4.09 4.09 0 014.09-4.09h5.218z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FileZip returns a string representing an "file-zip" Octicon.
func FileZip(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-zip%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.5 1.75a.25.25 0 01.25-.25h3a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h2.086a.25.25 0 01.177.073l2.914 2.914a.25.25 0 01.073.177v8.586a.25.25 0 01-.25.25h-.5a.75.75 0 000 1.5h.5A1.75 1.75 0 0014 13.25V4.664c0-.464-.184-.909-.513-1.237L10.573.513A1.75 1.75 0 009.336 0H3.75A1.75 1.75 0 002 1.75v11.5c0 .649.353 1.214.874 1.515a.75.75 0 10.752-1.298.25.25 0 01-.126-.217V1.75zM8.75 3a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM6 5.25a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5A.75.75 0 016 5.25zm2 1.5A.75.75 0 018.75 6h.5a.75.75 0 010 1.5h-.5A.75.75 0 018 6.75zm-1.25.75a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM8 9.75A.75.75 0 018.75 9h.5a.75.75 0 010 1.5h-.5A.75.75 0 018 9.75zm-.75.75a1.75 1.75 0 00-1.75 1.75v3c0 .414.336.75.75.75h2.5a.75.75 0 00.75-.75v-3a1.75 1.75 0 00-1.75-1.75h-.5zM7 12.25a.25.25 0 01.25-.25h.5a.25.25 0 01.25.25v2.25H7v-2.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-file-zip%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M5 2.5a.5.5 0 00-.5.5v18a.5.5 0 00.5.5h1.75a.75.75 0 010 1.5H5a2 2 0 01-2-2V3a2 2 0 012-2h9.982a2 2 0 011.414.586l4.018 4.018A2 2 0 0121 7.018V21a2 2 0 01-2 2h-2.75a.75.75 0 010-1.5H19a.5.5 0 00.5-.5V7.018a.5.5 0 00-.146-.354l-4.018-4.018a.5.5 0 00-.354-.146H5z\"></path><path d=\"M11.5 15.75a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm.75-3.75a.75.75 0 000 1.5h1a.75.75 0 000-1.5h-1zm-.75-2.25a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zM12.25 6a.75.75 0 000 1.5h1a.75.75 0 000-1.5h-1zm-.75-2.25a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zM9.75 13.5a.75.75 0 000 1.5h1a.75.75 0 000-1.5h-1zM9 11.25a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm.75-3.75a.75.75 0 000 1.5h1a.75.75 0 000-1.5h-1zM9 5.25a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1A.75.75 0 019 5.25z\"></path><path fill-rule=\"evenodd\" d=\"M11 17a2 2 0 00-2 2v4.25c0 .414.336.75.75.75h3.5a.75.75 0 00.75-.75V19a2 2 0 00-2-2h-1zm-.5 2a.5.5 0 01.5-.5h1a.5.5 0 01.5.5v3.5h-2V19z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Filter returns a string representing an "filter" Octicon.
func Filter(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-filter%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M.75 3a.75.75 0 000 1.5h14.5a.75.75 0 000-1.5H.75zM3 7.75A.75.75 0 013.75 7h8.5a.75.75 0 010 1.5h-8.5A.75.75 0 013 7.75zm3 4a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5h-2.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-filter%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M2.75 6a.75.75 0 000 1.5h18.5a.75.75 0 000-1.5H2.75zM6 11.75a.75.75 0 01.75-.75h10.5a.75.75 0 010 1.5H6.75a.75.75 0 01-.75-.75zm4 4.938a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5h-2.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Flame returns a string representing an "flame" Octicon.
func Flame(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-flame%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.998 14.5c2.832 0 5-1.98 5-4.5 0-1.463-.68-2.19-1.879-3.383l-.036-.037c-1.013-1.008-2.3-2.29-2.834-4.434-.322.256-.63.579-.864.953-.432.696-.621 1.58-.046 2.73.473.947.67 2.284-.278 3.232-.61.61-1.545.84-2.403.633a2.788 2.788 0 01-1.436-.874A3.21 3.21 0 003 10c0 2.53 2.164 4.5 4.998 4.5zM9.533.753C9.496.34 9.16.009 8.77.146 7.035.75 4.34 3.187 5.997 6.5c.344.689.285 1.218.003 1.5-.419.419-1.54.487-2.04-.832-.173-.454-.659-.762-1.035-.454C2.036 7.44 1.5 8.702 1.5 10c0 3.512 2.998 6 6.498 6s6.5-2.5 6.5-6c0-2.137-1.128-3.26-2.312-4.438-1.19-1.184-2.436-2.425-2.653-4.81z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-flame%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.185 21.5c4.059 0 7.065-2.84 7.065-6.75 0-2.337-1.093-3.489-2.678-5.158l-.021-.023c-1.44-1.517-3.139-3.351-3.649-6.557a6.14 6.14 0 00-1.911 1.76c-.787 1.144-1.147 2.633-.216 4.495.603 1.205.777 2.74-.277 3.794-.657.657-1.762 1.1-2.956.586-.752-.324-1.353-.955-1.838-1.79-.567.706-.954 1.74-.954 2.893 0 3.847 3.288 6.75 7.435 6.75zm2.08-19.873c-.017-.345-.296-.625-.632-.543-2.337.575-6.605 4.042-4.2 8.854.474.946.392 1.675.004 2.062-.64.64-1.874.684-2.875-1.815-.131-.327-.498-.509-.803-.334-1.547.888-2.509 2.86-2.509 4.899 0 4.829 4.122 8.25 8.935 8.25 4.812 0 8.565-3.438 8.565-8.25 0-2.939-1.466-4.482-3.006-6.102-1.61-1.694-3.479-3.476-3.479-7.021z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Fold returns a string representing an "fold" Octicon.
func Fold(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-fold%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M10.896 2H8.75V.75a.75.75 0 00-1.5 0V2H5.104a.25.25 0 00-.177.427l2.896 2.896a.25.25 0 00.354 0l2.896-2.896A.25.25 0 0010.896 2zM8.75 15.25a.75.75 0 01-1.5 0V14H5.104a.25.25 0 01-.177-.427l2.896-2.896a.25.25 0 01.354 0l2.896 2.896a.25.25 0 01-.177.427H8.75v1.25zm-6.5-6.5a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5zM6 8a.75.75 0 01-.75.75h-.5a.75.75 0 010-1.5h.5A.75.75 0 016 8zm2.25.75a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5zM12 8a.75.75 0 01-.75.75h-.5a.75.75 0 010-1.5h.5A.75.75 0 0112 8zm2.25.75a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-fold%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 15a.75.75 0 01.53.22l3.25 3.25a.75.75 0 11-1.06 1.06L12 16.81l-2.72 2.72a.75.75 0 01-1.06-1.06l3.25-3.25A.75.75 0 0112 15z\"></path><path fill-rule=\"evenodd\" d=\"M12 15.75a.75.75 0 01.75.75v5.75a.75.75 0 01-1.5 0V16.5a.75.75 0 01.75-.75zm.53-6.97a.75.75 0 01-1.06 0L8.22 5.53a.75.75 0 011.06-1.06L12 7.19l2.72-2.72a.75.75 0 111.06 1.06l-3.25 3.25z\"></path><path fill-rule=\"evenodd\" d=\"M12 8.5a.75.75 0 01-.75-.75v-6a.75.75 0 011.5 0v6a.75.75 0 01-.75.75zM10.75 12a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FoldDown returns a string representing an "fold-down" Octicon.
func FoldDown(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-fold-down%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8.177 14.323l2.896-2.896a.25.25 0 00-.177-.427H8.75V7.764a.75.75 0 10-1.5 0V11H5.104a.25.25 0 00-.177.427l2.896 2.896a.25.25 0 00.354 0zM2.25 5a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5zM6 4.25a.75.75 0 01-.75.75h-.5a.75.75 0 010-1.5h.5a.75.75 0 01.75.75zM8.25 5a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5zM12 4.25a.75.75 0 01-.75.75h-.5a.75.75 0 010-1.5h.5a.75.75 0 01.75.75zm2.25.75a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-fold-down%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 19a.75.75 0 01-.53-.22l-3.25-3.25a.75.75 0 111.06-1.06L12 17.19l2.72-2.72a.75.75 0 111.06 1.06l-3.25 3.25A.75.75 0 0112 19z\"></path><path fill-rule=\"evenodd\" d=\"M12 18a.75.75 0 01-.75-.75v-7.5a.75.75 0 011.5 0v7.5A.75.75 0 0112 18zM10.75 6a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1A.75.75 0 012.75 6zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1A.75.75 0 016.75 6zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// FoldUp returns a string representing an "fold-up" Octicon.
func FoldUp(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-fold-up%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M7.823 1.677L4.927 4.573A.25.25 0 005.104 5H7.25v3.236a.75.75 0 101.5 0V5h2.146a.25.25 0 00.177-.427L8.177 1.677a.25.25 0 00-.354 0zM13.75 11a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zm-3.75.75a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5a.75.75 0 01-.75-.75zM7.75 11a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM4 11.75a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5a.75.75 0 01-.75-.75zM1.75 11a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-fold-up%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.47 5.22a.75.75 0 011.06 0l3.25 3.25a.75.75 0 01-1.06 1.06L12 6.81 9.28 9.53a.75.75 0 01-1.06-1.06l3.25-3.25z\"></path><path fill-rule=\"evenodd\" d=\"M12 5.5a.75.75 0 01.75.75v8a.75.75 0 01-1.5 0v-8A.75.75 0 0112 5.5zM10.75 18a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Gear returns a string representing an "gear" Octicon.
func Gear(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-gear%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.429 1.525a6.593 6.593 0 011.142 0c.036.003.108.036.137.146l.289 1.105c.147.56.55.967.997 1.189.174.086.341.183.501.29.417.278.97.423 1.53.27l1.102-.303c.11-.03.175.016.195.046.219.31.41.641.573.989.014.031.022.11-.059.19l-.815.806c-.411.406-.562.957-.53 1.456a4.588 4.588 0 010 .582c-.032.499.119 1.05.53 1.456l.815.806c.08.08.073.159.059.19a6.494 6.494 0 01-.573.99c-.02.029-.086.074-.195.045l-1.103-.303c-.559-.153-1.112-.008-1.529.27-.16.107-.327.204-.5.29-.449.222-.851.628-.998 1.189l-.289 1.105c-.029.11-.101.143-.137.146a6.613 6.613 0 01-1.142 0c-.036-.003-.108-.037-.137-.146l-.289-1.105c-.147-.56-.55-.967-.997-1.189a4.502 4.502 0 01-.501-.29c-.417-.278-.97-.423-1.53-.27l-1.102.303c-.11.03-.175-.016-.195-.046a6.492 6.492 0 01-.573-.989c-.014-.031-.022-.11.059-.19l.815-.806c.411-.406.562-.957.53-1.456a4.587 4.587 0 010-.582c.032-.499-.119-1.05-.53-1.456l-.815-.806c-.08-.08-.073-.159-.059-.19a6.44 6.44 0 01.573-.99c.02-.029.086-.075.195-.045l1.103.303c.559.153 1.112.008 1.529-.27.16-.107.327-.204.5-.29.449-.222.851-.628.998-1.189l.289-1.105c.029-.11.101-.143.137-.146zM8 0c-.236 0-.47.01-.701.03-.743.065-1.29.615-1.458 1.261l-.29 1.106c-.017.066-.078.158-.211.224a5.994 5.994 0 00-.668.386c-.123.082-.233.09-.3.071L3.27 2.776c-.644-.177-1.392.02-1.82.63a7.977 7.977 0 00-.704 1.217c-.315.675-.111 1.422.363 1.891l.815.806c.05.048.098.147.088.294a6.084 6.084 0 000 .772c.01.147-.038.246-.088.294l-.815.806c-.474.469-.678 1.216-.363 1.891.2.428.436.835.704 1.218.428.609 1.176.806 1.82.63l1.103-.303c.066-.019.176-.011.299.071.213.143.436.272.668.386.133.066.194.158.212.224l.289 1.106c.169.646.715 1.196 1.458 1.26a8.094 8.094 0 001.402 0c.743-.064 1.29-.614 1.458-1.26l.29-1.106c.017-.066.078-.158.211-.224a5.98 5.98 0 00.668-.386c.123-.082.233-.09.3-.071l1.102.302c.644.177 1.392-.02 1.82-.63.268-.382.505-.789.704-1.217.315-.675.111-1.422-.364-1.891l-.814-.806c-.05-.048-.098-.147-.088-.294a6.1 6.1 0 000-.772c-.01-.147.039-.246.088-.294l.814-.806c.475-.469.679-1.216.364-1.891a7.992 7.992 0 00-.704-1.218c-.428-.609-1.176-.806-1.82-.63l-1.103.303c-.066.019-.176.011-.299-.071a5.991 5.991 0 00-.668-.386c-.133-.066-.194-.158-.212-.224L10.16 1.29C9.99.645 9.444.095 8.701.031A8.094 8.094 0 008 0zm1.5 8a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM11 8a3 3 0 11-6 0 3 3 0 016 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-gear%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M16 12a4 4 0 11-8 0 4 4 0 018 0zm-1.5 0a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z\"></path><path fill-rule=\"evenodd\" d=\"M12 1c-.268 0-.534.01-.797.028-.763.055-1.345.617-1.512 1.304l-.352 1.45c-.02.078-.09.172-.225.22a8.45 8.45 0 00-.728.303c-.13.06-.246.044-.315.002l-1.274-.776c-.604-.368-1.412-.354-1.99.147-.403.348-.78.726-1.129 1.128-.5.579-.515 1.387-.147 1.99l.776 1.275c.042.069.059.185-.002.315-.112.237-.213.48-.302.728-.05.135-.143.206-.221.225l-1.45.352c-.687.167-1.249.749-1.304 1.512a11.149 11.149 0 000 1.594c.055.763.617 1.345 1.304 1.512l1.45.352c.078.02.172.09.22.225.09.248.191.491.303.729.06.129.044.245.002.314l-.776 1.274c-.368.604-.354 1.412.147 1.99.348.403.726.78 1.128 1.129.579.5 1.387.515 1.99.147l1.275-.776c.069-.042.185-.059.315.002.237.112.48.213.728.302.135.05.206.143.225.221l.352 1.45c.167.687.749 1.249 1.512 1.303a11.125 11.125 0 001.594 0c.763-.054 1.345-.616 1.512-1.303l.352-1.45c.02-.078.09-.172.225-.22.248-.09.491-.191.729-.303.129-.06.245-.044.314-.002l1.274.776c.604.368 1.412.354 1.99-.147.403-.348.78-.726 1.129-1.128.5-.579.515-1.387.147-1.99l-.776-1.275c-.042-.069-.059-.185.002-.315.112-.237.213-.48.302-.728.05-.135.143-.206.221-.225l1.45-.352c.687-.167 1.249-.749 1.303-1.512a11.125 11.125 0 000-1.594c-.054-.763-.616-1.345-1.303-1.512l-1.45-.352c-.078-.02-.172-.09-.22-.225a8.469 8.469 0 00-.303-.728c-.06-.13-.044-.246-.002-.315l.776-1.274c.368-.604.354-1.412-.147-1.99-.348-.403-.726-.78-1.128-1.129-.579-.5-1.387-.515-1.99-.147l-1.275.776c-.069.042-.185.059-.315-.002a8.465 8.465 0 00-.728-.302c-.135-.05-.206-.143-.225-.221l-.352-1.45c-.167-.687-.749-1.249-1.512-1.304A11.149 11.149 0 0012 1zm-.69 1.525a9.648 9.648 0 011.38 0c.055.004.135.05.162.16l.351 1.45c.153.628.626 1.08 1.173 1.278.205.074.405.157.6.249a1.832 1.832 0 001.733-.074l1.275-.776c.097-.06.186-.036.228 0 .348.302.674.628.976.976.036.042.06.13 0 .228l-.776 1.274a1.832 1.832 0 00-.074 1.734c.092.195.175.395.248.6.198.547.652 1.02 1.278 1.172l1.45.353c.111.026.157.106.161.161a9.653 9.653 0 010 1.38c-.004.055-.05.135-.16.162l-1.45.351a1.833 1.833 0 00-1.278 1.173 6.926 6.926 0 01-.25.6 1.832 1.832 0 00.075 1.733l.776 1.275c.06.097.036.186 0 .228a9.555 9.555 0 01-.976.976c-.042.036-.13.06-.228 0l-1.275-.776a1.832 1.832 0 00-1.733-.074 6.926 6.926 0 01-.6.248 1.833 1.833 0 00-1.172 1.278l-.353 1.45c-.026.111-.106.157-.161.161a9.653 9.653 0 01-1.38 0c-.055-.004-.135-.05-.162-.16l-.351-1.45a1.833 1.833 0 00-1.173-1.278 6.928 6.928 0 01-.6-.25 1.832 1.832 0 00-1.734.075l-1.274.776c-.097.06-.186.036-.228 0a9.56 9.56 0 01-.976-.976c-.036-.042-.06-.13 0-.228l.776-1.275a1.832 1.832 0 00.074-1.733 6.948 6.948 0 01-.249-.6 1.833 1.833 0 00-1.277-1.172l-1.45-.353c-.111-.026-.157-.106-.161-.161a9.648 9.648 0 010-1.38c.004-.055.05-.135.16-.162l1.45-.351a1.833 1.833 0 001.278-1.173 6.95 6.95 0 01.249-.6 1.832 1.832 0 00-.074-1.734l-.776-1.274c-.06-.097-.036-.186 0-.228.302-.348.628-.674.976-.976.042-.036.13-.06.228 0l1.274.776a1.832 1.832 0 001.734.074 6.95 6.95 0 01.6-.249 1.833 1.833 0 001.172-1.277l.353-1.45c.026-.111.106-.157.161-.161z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Gift returns a string representing an "gift" Octicon.
func Gift(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-gift%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 1.5a1.25 1.25 0 100 2.5h2.309c-.233-.818-.542-1.401-.878-1.793-.43-.502-.915-.707-1.431-.707zM2 2.75c0 .45.108.875.3 1.25h-.55A1.75 1.75 0 000 5.75v2c0 .698.409 1.3 1 1.582v4.918c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0015 14.25V9.332c.591-.281 1-.884 1-1.582v-2A1.75 1.75 0 0014.25 4h-.55a2.75 2.75 0 00-2.45-4c-.984 0-1.874.42-2.57 1.23A5.086 5.086 0 008 2.274a5.086 5.086 0 00-.68-1.042C6.623.42 5.733 0 4.75 0A2.75 2.75 0 002 2.75zM8.941 4h2.309a1.25 1.25 0 100-2.5c-.516 0-1 .205-1.43.707-.337.392-.646.975-.879 1.793zm-1.84 1.5H1.75a.25.25 0 00-.25.25v2c0 .138.112.25.25.25h5.5V5.5h-.149zm1.649 0V8h5.5a.25.25 0 00.25-.25v-2a.25.25 0 00-.25-.25h-5.5zm0 4h4.75v4.75a.25.25 0 01-.25.25h-4.5v-5zm-1.5 0v5h-4.5a.25.25 0 01-.25-.25V9.5h4.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-gift%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 3.75c0 .844.279 1.623.75 2.25H2.75A1.75 1.75 0 001 7.75v2.5c0 .698.409 1.3 1 1.582v8.418c0 .966.784 1.75 1.75 1.75h16.5A1.75 1.75 0 0022 20.25v-8.418c.591-.281 1-.884 1-1.582v-2.5A1.75 1.75 0 0021.25 6H19.5a3.75 3.75 0 00-3-6c-1.456 0-3.436.901-4.5 3.11C10.936.901 8.955 0 7.5 0a3.75 3.75 0 00-3.75 3.75zM11.22 6c-.287-3.493-2.57-4.5-3.72-4.5a2.25 2.25 0 000 4.5h3.72zm9.28 6v8.25a.25.25 0 01-.25.25h-7.5V12h7.75zm-9.25 8.5V12H3.5v8.25c0 .138.112.25.25.25h7.5zm10-10a.25.25 0 00.25-.25v-2.5a.25.25 0 00-.25-.25h-8.5v3h8.5zm-18.5 0h8.5v-3h-8.5a.25.25 0 00-.25.25v2.5c0 .138.112.25.25.25zm16-6.75A2.25 2.25 0 0116.5 6h-3.72c.287-3.493 2.57-4.5 3.72-4.5a2.25 2.25 0 012.25 2.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// GitBranch returns a string representing an "git-branch" Octicon.
func GitBranch(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-branch%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.75 2.5a.75.75 0 100 1.5.75.75 0 000-1.5zm-2.25.75a2.25 2.25 0 113 2.122V6A2.5 2.5 0 0110 8.5H6a1 1 0 00-1 1v1.128a2.251 2.251 0 11-1.5 0V5.372a2.25 2.25 0 111.5 0v1.836A2.492 2.492 0 016 7h4a1 1 0 001-1v-.628A2.25 2.25 0 019.5 3.25zM4.25 12a.75.75 0 100 1.5.75.75 0 000-1.5zM3.5 3.25a.75.75 0 111.5 0 .75.75 0 01-1.5 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-branch%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.75 21a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM2.5 19.25a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0zM5.75 6.5a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM2.5 4.75a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0zM18.25 6.5a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM15 4.75a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0z\"></path><path fill-rule=\"evenodd\" d=\"M5.75 16.75A.75.75 0 006.5 16V8A.75.75 0 005 8v8c0 .414.336.75.75.75z\"></path><path fill-rule=\"evenodd\" d=\"M17.5 8.75v-1H19v1a3.75 3.75 0 01-3.75 3.75h-7a1.75 1.75 0 00-1.75 1.75H5A3.25 3.25 0 018.25 11h7a2.25 2.25 0 002.25-2.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// GitCommit returns a string representing an "git-commit" Octicon.
func GitCommit(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-commit%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.5 7.75a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0zm1.43.75a4.002 4.002 0 01-7.86 0H.75a.75.75 0 110-1.5h3.32a4.001 4.001 0 017.86 0h3.32a.75.75 0 110 1.5h-3.32z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-commit%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M15.5 11.75a3.5 3.5 0 11-7 0 3.5 3.5 0 017 0zm1.444-.75a5.001 5.001 0 00-9.888 0H2.75a.75.75 0 100 1.5h4.306a5.001 5.001 0 009.888 0h4.306a.75.75 0 100-1.5h-4.306z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// GitCompare returns a string representing an "git-compare" Octicon.
func GitCompare(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-compare%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.573.677L7.177 3.073a.25.25 0 000 .354l2.396 2.396A.25.25 0 0010 5.646V4h1a1 1 0 011 1v5.628a2.251 2.251 0 101.5 0V5A2.5 2.5 0 0011 2.5h-1V.854a.25.25 0 00-.427-.177zM6 12v-1.646a.25.25 0 01.427-.177l2.396 2.396a.25.25 0 010 .354l-2.396 2.396A.25.25 0 016 15.146V13.5H5A2.5 2.5 0 012.5 11V5.372a2.25 2.25 0 111.5 0V11a1 1 0 001 1h1zm6.75 0a.75.75 0 100 1.5.75.75 0 000-1.5zM4 3.25a.75.75 0 11-1.5 0 .75.75 0 011.5 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-compare%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M19.75 17.5a1.75 1.75 0 100 3.5 1.75 1.75 0 000-3.5zm-3.25 1.75a3.25 3.25 0 116.5 0 3.25 3.25 0 01-6.5 0z\"></path><path fill-rule=\"evenodd\" d=\"M13.905 1.72a.75.75 0 010 1.06L12.685 4h4.065a3.75 3.75 0 013.75 3.75v8.75a.75.75 0 01-1.5 0V7.75a2.25 2.25 0 00-2.25-2.25h-4.064l1.22 1.22a.75.75 0 01-1.061 1.06l-2.5-2.5a.75.75 0 010-1.06l2.5-2.5a.75.75 0 011.06 0zM4.25 6.5a1.75 1.75 0 100-3.5 1.75 1.75 0 000 3.5zM7.5 4.75a3.25 3.25 0 11-6.5 0 3.25 3.25 0 016.5 0z\"></path><path fill-rule=\"evenodd\" d=\"M10.095 22.28a.75.75 0 010-1.06l1.22-1.22H7.25a3.75 3.75 0 01-3.75-3.75V7.5a.75.75 0 011.5 0v8.75a2.25 2.25 0 002.25 2.25h4.064l-1.22-1.22a.75.75 0 111.061-1.06l2.5 2.5a.75.75 0 010 1.06l-2.5 2.5a.75.75 0 01-1.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// GitMerge returns a string representing an "git-merge" Octicon.
func GitMerge(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-merge%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5 3.254V3.25v.005a.75.75 0 110-.005v.004zm.45 1.9a2.25 2.25 0 10-1.95.218v5.256a2.25 2.25 0 101.5 0V7.123A5.735 5.735 0 009.25 9h1.378a2.251 2.251 0 100-1.5H9.25a4.25 4.25 0 01-3.8-2.346zM12.75 9a.75.75 0 100-1.5.75.75 0 000 1.5zm-8.5 4.5a.75.75 0 100-1.5.75.75 0 000 1.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-merge%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.75 21a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM2.5 19.25a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0zM5.75 6.5a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM2.5 4.75a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0zM18.25 15a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM15 13.25a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0z\"></path><path fill-rule=\"evenodd\" d=\"M6.5 7.25c0 2.9 2.35 5.25 5.25 5.25h4.5V14h-4.5A6.75 6.75 0 015 7.25h1.5z\"></path><path fill-rule=\"evenodd\" d=\"M5.75 16.75A.75.75 0 006.5 16V8A.75.75 0 005 8v8c0 .414.336.75.75.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// GitPullRequest returns a string representing an "git-pull-request" Octicon.
func GitPullRequest(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-pull-request%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.177 3.073L9.573.677A.25.25 0 0110 .854v4.792a.25.25 0 01-.427.177L7.177 3.427a.25.25 0 010-.354zM3.75 2.5a.75.75 0 100 1.5.75.75 0 000-1.5zm-2.25.75a2.25 2.25 0 113 2.122v5.256a2.251 2.251 0 11-1.5 0V5.372A2.25 2.25 0 011.5 3.25zM11 2.5h-1V4h1a1 1 0 011 1v5.628a2.251 2.251 0 101.5 0V5A2.5 2.5 0 0011 2.5zm1 10.25a.75.75 0 111.5 0 .75.75 0 01-1.5 0zM3.75 12a.75.75 0 100 1.5.75.75 0 000-1.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-pull-request%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 3a1.75 1.75 0 100 3.5 1.75 1.75 0 000-3.5zM1.5 4.75a3.25 3.25 0 116.5 0 3.25 3.25 0 01-6.5 0zM4.75 17.5a1.75 1.75 0 100 3.5 1.75 1.75 0 000-3.5zM1.5 19.25a3.25 3.25 0 116.5 0 3.25 3.25 0 01-6.5 0zm17.75-1.75a1.75 1.75 0 100 3.5 1.75 1.75 0 000-3.5zM16 19.25a3.25 3.25 0 116.5 0 3.25 3.25 0 01-6.5 0z\"></path><path fill-rule=\"evenodd\" d=\"M4.75 7.25A.75.75 0 015.5 8v8A.75.75 0 014 16V8a.75.75 0 01.75-.75zm8.655-5.53a.75.75 0 010 1.06L12.185 4h4.065A3.75 3.75 0 0120 7.75v8.75a.75.75 0 01-1.5 0V7.75a2.25 2.25 0 00-2.25-2.25h-4.064l1.22 1.22a.75.75 0 01-1.061 1.06l-2.5-2.5a.75.75 0 010-1.06l2.5-2.5a.75.75 0 011.06 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// GitPullRequestClosed returns a string representing an "git-pull-request-closed" Octicon.
func GitPullRequestClosed(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-pull-request-closed%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.72 1.227a.75.75 0 011.06 0l.97.97.97-.97a.75.75 0 111.06 1.061l-.97.97.97.97a.75.75 0 01-1.06 1.06l-.97-.97-.97.97a.75.75 0 11-1.06-1.06l.97-.97-.97-.97a.75.75 0 010-1.06zM12.75 6.5a.75.75 0 00-.75.75v3.378a2.251 2.251 0 101.5 0V7.25a.75.75 0 00-.75-.75zm0 5.5a.75.75 0 100 1.5.75.75 0 000-1.5zM2.5 3.25a.75.75 0 111.5 0 .75.75 0 01-1.5 0zM3.25 1a2.25 2.25 0 00-.75 4.372v5.256a2.251 2.251 0 101.5 0V5.372A2.25 2.25 0 003.25 1zm0 11a.75.75 0 100 1.5.75.75 0 000-1.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-pull-request-closed%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M22.266 2.711a.75.75 0 10-1.061-1.06l-1.983 1.983-1.984-1.983a.75.75 0 10-1.06 1.06l1.983 1.983-1.983 1.984a.75.75 0 001.06 1.06l1.984-1.983 1.983 1.983a.75.75 0 001.06-1.06l-1.983-1.984 1.984-1.983z\"></path><path fill-rule=\"evenodd\" d=\"M4.75 1.5a3.25 3.25 0 00-.745 6.414A.758.758 0 004 8v8a.81.81 0 00.005.086A3.251 3.251 0 004.75 22.5a3.25 3.25 0 00.745-6.414A.758.758 0 005.5 16V8a.758.758 0 00-.005-.086A3.251 3.251 0 004.75 1.5zM3 4.75a1.75 1.75 0 113.5 0 1.75 1.75 0 01-3.5 0zm0 14.5a1.75 1.75 0 113.5 0 1.75 1.75 0 01-3.5 0zm13 0a3.251 3.251 0 012.5-3.163V9.625a.75.75 0 011.5 0v6.462a3.251 3.251 0 01-.75 6.413A3.25 3.25 0 0116 19.25zm3.25-1.75a1.75 1.75 0 100 3.5 1.75 1.75 0 000-3.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// GitPullRequestDraft returns a string representing an "git-pull-request-draft" Octicon.
func GitPullRequestDraft(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-pull-request-draft%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.5 3.25a.75.75 0 111.5 0 .75.75 0 01-1.5 0zM3.25 1a2.25 2.25 0 00-.75 4.372v5.256a2.251 2.251 0 101.5 0V5.372A2.25 2.25 0 003.25 1zm0 11a.75.75 0 100 1.5.75.75 0 000-1.5zm9.5 3a2.25 2.25 0 100-4.5 2.25 2.25 0 000 4.5zm0-3a.75.75 0 100 1.5.75.75 0 000-1.5z\"></path><path d=\"M14 7.5a1.25 1.25 0 11-2.5 0 1.25 1.25 0 012.5 0zm0-4.25a1.25 1.25 0 11-2.5 0 1.25 1.25 0 012.5 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-git-pull-request-draft%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 1.5a3.25 3.25 0 00-.745 6.414A.758.758 0 004 8v8a.81.81 0 00.005.086A3.251 3.251 0 004.75 22.5a3.25 3.25 0 00.745-6.414A.757.757 0 005.5 16V8a.758.758 0 00-.005-.086A3.251 3.251 0 004.75 1.5zM3 4.75a1.75 1.75 0 113.5 0 1.75 1.75 0 01-3.5 0zm0 14.5a1.75 1.75 0 113.5 0 1.75 1.75 0 01-3.5 0zm13 0a3.25 3.25 0 116.5 0 3.25 3.25 0 01-6.5 0zm3.25-1.75a1.75 1.75 0 100 3.5 1.75 1.75 0 000-3.5z\"></path><path d=\"M19.25 6a1.75 1.75 0 100-3.5 1.75 1.75 0 000 3.5zM21 11.25a1.75 1.75 0 11-3.5 0 1.75 1.75 0 013.5 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Globe returns a string representing an "globe" Octicon.
func Globe(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-globe%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.543 7.25h2.733c.144-2.074.866-3.756 1.58-4.948.12-.197.237-.381.353-.552a6.506 6.506 0 00-4.666 5.5zm2.733 1.5H1.543a6.506 6.506 0 004.666 5.5 11.13 11.13 0 01-.352-.552c-.715-1.192-1.437-2.874-1.581-4.948zm1.504 0h4.44a9.637 9.637 0 01-1.363 4.177c-.306.51-.612.919-.857 1.215a9.978 9.978 0 01-.857-1.215A9.637 9.637 0 015.78 8.75zm4.44-1.5H5.78a9.637 9.637 0 011.363-4.177c.306-.51.612-.919.857-1.215.245.296.55.705.857 1.215A9.638 9.638 0 0110.22 7.25zm1.504 1.5c-.144 2.074-.866 3.756-1.58 4.948-.12.197-.237.381-.353.552a6.506 6.506 0 004.666-5.5h-2.733zm2.733-1.5h-2.733c-.144-2.074-.866-3.756-1.58-4.948a11.738 11.738 0 00-.353-.552 6.506 6.506 0 014.666 5.5zM8 0a8 8 0 100 16A8 8 0 008 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-globe%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.513 11.5h4.745c.1-3.037 1.1-5.49 2.093-7.204.39-.672.78-1.233 1.119-1.673C6.11 3.329 2.746 7 2.513 11.5zm4.77 1.5H2.552a9.505 9.505 0 007.918 8.377 15.698 15.698 0 01-1.119-1.673C8.413 18.085 7.47 15.807 7.283 13zm1.504 0h6.426c-.183 2.48-1.02 4.5-1.862 5.951-.476.82-.95 1.455-1.304 1.88L12 20.89l-.047-.057a13.888 13.888 0 01-1.304-1.88C9.807 17.5 8.969 15.478 8.787 13zm6.454-1.5H8.759c.1-2.708.992-4.904 1.89-6.451.476-.82.95-1.455 1.304-1.88L12 3.11l.047.057c.353.426.828 1.06 1.304 1.88.898 1.548 1.79 3.744 1.89 6.452zm1.476 1.5c-.186 2.807-1.13 5.085-2.068 6.704-.39.672-.78 1.233-1.118 1.673A9.505 9.505 0 0021.447 13h-4.731zm4.77-1.5h-4.745c-.1-3.037-1.1-5.49-2.093-7.204-.39-.672-.78-1.233-1.119-1.673 4.36.706 7.724 4.377 7.957 8.877z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Grabber returns a string representing an "grabber" Octicon.
func Grabber(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-grabber%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10 13a1 1 0 100-2 1 1 0 000 2zm-4 0a1 1 0 100-2 1 1 0 000 2zm1-5a1 1 0 11-2 0 1 1 0 012 0zm3 1a1 1 0 100-2 1 1 0 000 2zm1-5a1 1 0 11-2 0 1 1 0 012 0zM6 5a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-grabber%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M15 18a1 1 0 100-2 1 1 0 000 2zm1-6a1 1 0 11-2 0 1 1 0 012 0zm-7 6a1 1 0 100-2 1 1 0 000 2zm0-5a1 1 0 100-2 1 1 0 000 2zm7-6a1 1 0 11-2 0 1 1 0 012 0zM9 8a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Graph returns a string representing an "graph" Octicon.
func Graph(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-graph%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 1.75a.75.75 0 00-1.5 0v12.5c0 .414.336.75.75.75h14.5a.75.75 0 000-1.5H1.5V1.75zm14.28 2.53a.75.75 0 00-1.06-1.06L10 7.94 7.53 5.47a.75.75 0 00-1.06 0L3.22 8.72a.75.75 0 001.06 1.06L7 7.06l2.47 2.47a.75.75 0 001.06 0l5.25-5.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-graph%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M2.5 2.75a.75.75 0 00-1.5 0v18.5c0 .414.336.75.75.75H20a.75.75 0 000-1.5H2.5V2.75z\"></path><path d=\"M22.28 7.78a.75.75 0 00-1.06-1.06l-5.72 5.72-3.72-3.72a.75.75 0 00-1.06 0l-6 6a.75.75 0 101.06 1.06l5.47-5.47 3.72 3.72a.75.75 0 001.06 0l6.25-6.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Hash returns a string representing an "hash" Octicon.
func Hash(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-hash%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.368 1.01a.75.75 0 01.623.859L6.57 4.5h3.98l.46-2.868a.75.75 0 011.48.237L12.07 4.5h2.18a.75.75 0 010 1.5h-2.42l-.64 4h2.56a.75.75 0 010 1.5h-2.8l-.46 2.869a.75.75 0 01-1.48-.237l.42-2.632H5.45l-.46 2.869a.75.75 0 01-1.48-.237l.42-2.632H1.75a.75.75 0 010-1.5h2.42l.64-4H2.25a.75.75 0 010-1.5h2.8l.46-2.868a.75.75 0 01.858-.622zM9.67 10l.64-4H6.33l-.64 4h3.98z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-hash%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.618 1.76a.75.75 0 01.623.859L9.46 7.5h6.48l.82-5.118a.75.75 0 011.48.237L17.46 7.5h3.79a.75.75 0 010 1.5h-4.03l-.96 6h3.99a.75.75 0 110 1.5h-4.23l-.78 4.869a.75.75 0 01-1.48-.237l.74-4.632H8.02l-.78 4.869a.75.75 0 01-1.48-.237L6.5 16.5H2.745a.75.75 0 010-1.5H6.74l.96-6H3.75a.75.75 0 010-1.5h4.19l.82-5.118a.75.75 0 01.858-.622zM14.741 15l.96-6H9.22l-.96 6h6.48z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Heading returns a string representing an "heading" Octicon.
func Heading(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-heading%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 2a.75.75 0 01.75.75V7h7V2.75a.75.75 0 011.5 0v10.5a.75.75 0 01-1.5 0V8.5h-7v4.75a.75.75 0 01-1.5 0V2.75A.75.75 0 013.75 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-heading%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.25 4a.75.75 0 01.75.75V11h10V4.75a.75.75 0 011.5 0v14.5a.75.75 0 01-1.5 0V12.5H7v6.75a.75.75 0 01-1.5 0V4.75A.75.75 0 016.25 4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Heart returns a string representing an "heart" Octicon.
func Heart(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-heart%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.25 2.5c-1.336 0-2.75 1.164-2.75 3 0 2.15 1.58 4.144 3.365 5.682A20.565 20.565 0 008 13.393a20.561 20.561 0 003.135-2.211C12.92 9.644 14.5 7.65 14.5 5.5c0-1.836-1.414-3-2.75-3-1.373 0-2.609.986-3.029 2.456a.75.75 0 01-1.442 0C6.859 3.486 5.623 2.5 4.25 2.5zM8 14.25l-.345.666-.002-.001-.006-.003-.018-.01a7.643 7.643 0 01-.31-.17 22.075 22.075 0 01-3.434-2.414C2.045 10.731 0 8.35 0 5.5 0 2.836 2.086 1 4.25 1 5.797 1 7.153 1.802 8 3.02 8.847 1.802 10.203 1 11.75 1 13.914 1 16 2.836 16 5.5c0 2.85-2.045 5.231-3.885 6.818a22.08 22.08 0 01-3.744 2.584l-.018.01-.006.003h-.002L8 14.25zm0 0l.345.666a.752.752 0 01-.69 0L8 14.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-heart%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.736 4C4.657 4 2.5 5.88 2.5 8.514c0 3.107 2.324 5.96 4.861 8.12a29.66 29.66 0 004.566 3.175l.073.041.073-.04c.271-.153.661-.38 1.13-.674.94-.588 2.19-1.441 3.436-2.502 2.537-2.16 4.861-5.013 4.861-8.12C21.5 5.88 19.343 4 17.264 4c-2.106 0-3.801 1.389-4.553 3.643a.75.75 0 01-1.422 0C10.537 5.389 8.841 4 6.736 4zM12 20.703l.343.667a.75.75 0 01-.686 0l.343-.667zM1 8.513C1 5.053 3.829 2.5 6.736 2.5 9.03 2.5 10.881 3.726 12 5.605 13.12 3.726 14.97 2.5 17.264 2.5 20.17 2.5 23 5.052 23 8.514c0 3.818-2.801 7.06-5.389 9.262a31.146 31.146 0 01-5.233 3.576l-.025.013-.007.003-.002.001-.344-.666-.343.667-.003-.002-.007-.003-.025-.013A29.308 29.308 0 0110 20.408a31.147 31.147 0 01-3.611-2.632C3.8 15.573 1 12.332 1 8.514z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// HeartFill returns a string representing an "heart-fill" Octicon.
func HeartFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-heart-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.655 14.916L8 14.25l.345.666a.752.752 0 01-.69 0zm0 0L8 14.25l.345.666.002-.001.006-.003.018-.01a7.643 7.643 0 00.31-.17 22.08 22.08 0 003.433-2.414C13.956 10.731 16 8.35 16 5.5 16 2.836 13.914 1 11.75 1 10.203 1 8.847 1.802 8 3.02 7.153 1.802 5.797 1 4.25 1 2.086 1 0 2.836 0 5.5c0 2.85 2.045 5.231 3.885 6.818a22.075 22.075 0 003.744 2.584l.018.01.006.003h.002z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-heart-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M14 20.408c-.492.308-.903.546-1.192.709-.153.086-.308.17-.463.252h-.002a.75.75 0 01-.686 0 16.709 16.709 0 01-.465-.252 31.147 31.147 0 01-4.803-3.34C3.8 15.572 1 12.331 1 8.513 1 5.052 3.829 2.5 6.736 2.5 9.03 2.5 10.881 3.726 12 5.605 13.12 3.726 14.97 2.5 17.264 2.5 20.17 2.5 23 5.052 23 8.514c0 3.818-2.801 7.06-5.389 9.262A31.146 31.146 0 0114 20.408z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// History returns a string representing an "history" Octicon.
func History(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-history%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.643 3.143L.427 1.927A.25.25 0 000 2.104V5.75c0 .138.112.25.25.25h3.646a.25.25 0 00.177-.427L2.715 4.215a6.5 6.5 0 11-1.18 4.458.75.75 0 10-1.493.154 8.001 8.001 0 101.6-5.684zM7.75 4a.75.75 0 01.75.75v2.992l2.028.812a.75.75 0 01-.557 1.392l-2.5-1A.75.75 0 017 8.25v-3.5A.75.75 0 017.75 4z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-history%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M11.998 2.5A9.503 9.503 0 003.378 8H5.75a.75.75 0 010 1.5H2a1 1 0 01-1-1V4.75a.75.75 0 011.5 0v1.697A10.997 10.997 0 0111.998 1C18.074 1 23 5.925 23 12s-4.926 11-11.002 11C6.014 23 1.146 18.223 1 12.275a.75.75 0 011.5-.037 9.5 9.5 0 009.498 9.262c5.248 0 9.502-4.253 9.502-9.5s-4.254-9.5-9.502-9.5z\"></path><path d=\"M12.5 7.25a.75.75 0 00-1.5 0v5.5c0 .27.144.518.378.651l3.5 2a.75.75 0 00.744-1.302L12.5 12.315V7.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Home returns a string representing an "home" Octicon.
func Home(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-home%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.156 1.835a.25.25 0 00-.312 0l-5.25 4.2a.25.25 0 00-.094.196v7.019c0 .138.112.25.25.25H5.5V8.25a.75.75 0 01.75-.75h3.5a.75.75 0 01.75.75v5.25h2.75a.25.25 0 00.25-.25V6.23a.25.25 0 00-.094-.195l-5.25-4.2zM6.906.664a1.75 1.75 0 012.187 0l5.25 4.2c.415.332.657.835.657 1.367v7.019A1.75 1.75 0 0113.25 15h-3.5a.75.75 0 01-.75-.75V9H7v5.25a.75.75 0 01-.75.75h-3.5A1.75 1.75 0 011 13.25V6.23c0-.531.242-1.034.657-1.366l5.25-4.2h-.001z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-home%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.03 2.59a1.5 1.5 0 011.94 0l7.5 6.363a1.5 1.5 0 01.53 1.144V19.5a1.5 1.5 0 01-1.5 1.5h-5.75a.75.75 0 01-.75-.75V14h-2v6.25a.75.75 0 01-.75.75H4.5A1.5 1.5 0 013 19.5v-9.403c0-.44.194-.859.53-1.144l7.5-6.363zM12 3.734l-7.5 6.363V19.5h5v-6.25a.75.75 0 01.75-.75h3.5a.75.75 0 01.75.75v6.25h5v-9.403L12 3.734z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// HomeFill returns a string representing an "home-fill" Octicon.
func HomeFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-home-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12.97 2.59a1.5 1.5 0 00-1.94 0l-7.5 6.363A1.5 1.5 0 003 10.097V19.5A1.5 1.5 0 004.5 21h4.75a.75.75 0 00.75-.75V14h4v6.25c0 .414.336.75.75.75h4.75a1.5 1.5 0 001.5-1.5v-9.403a1.5 1.5 0 00-.53-1.144l-7.5-6.363z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// HorizontalRule returns a string representing an "horizontal-rule" Octicon.
func HorizontalRule(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-horizontal-rule%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 7.75A.75.75 0 01.75 7h14.5a.75.75 0 010 1.5H.75A.75.75 0 010 7.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-horizontal-rule%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 12.75a.75.75 0 01.75-.75h18.5a.75.75 0 010 1.5H2.75a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Hourglass returns a string representing an "hourglass" Octicon.
func Hourglass(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-hourglass%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 1a.75.75 0 000 1.5h.75v1.25a4.75 4.75 0 001.9 3.8l.333.25c.134.1.134.3 0 .4l-.333.25a4.75 4.75 0 00-1.9 3.8v1.25h-.75a.75.75 0 000 1.5h10.5a.75.75 0 000-1.5h-.75v-1.25a4.75 4.75 0 00-1.9-3.8l-.333-.25a.25.25 0 010-.4l.333-.25a4.75 4.75 0 001.9-3.8V2.5h.75a.75.75 0 000-1.5H2.75zM11 2.5H5v1.25a3.25 3.25 0 001.3 2.6l.333.25c.934.7.934 2.1 0 2.8l-.333.25a3.25 3.25 0 00-1.3 2.6v1.25h6v-1.25a3.25 3.25 0 00-1.3-2.6l-.333-.25a1.75 1.75 0 010-2.8l.333-.25a3.25 3.25 0 001.3-2.6V2.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-hourglass%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 2a.75.75 0 000 1.5h.75v2.982a4.75 4.75 0 002.215 4.017l2.044 1.29a.25.25 0 010 .422l-2.044 1.29A4.75 4.75 0 005.5 17.518V20.5h-.75a.75.75 0 000 1.5h14.5a.75.75 0 000-1.5h-.75v-2.982a4.75 4.75 0 00-2.215-4.017l-2.044-1.29a.25.25 0 010-.422l2.044-1.29A4.75 4.75 0 0018.5 6.482V3.5h.75a.75.75 0 000-1.5H4.75zM17 3.5H7v2.982A3.25 3.25 0 008.516 9.23l2.044 1.29a1.75 1.75 0 010 2.96l-2.044 1.29A3.25 3.25 0 007 17.518V20.5h10v-2.982a3.25 3.25 0 00-1.516-2.748l-2.044-1.29a1.75 1.75 0 010-2.96l2.044-1.29A3.25 3.25 0 0017 6.482V3.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Hubot returns a string representing an "hubot" Octicon.
func Hubot(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-hubot%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 8a8 8 0 1116 0v5.25a.75.75 0 01-1.5 0V8a6.5 6.5 0 10-13 0v5.25a.75.75 0 01-1.5 0V8zm5.5 4.25a.75.75 0 01.75-.75h3.5a.75.75 0 010 1.5h-3.5a.75.75 0 01-.75-.75zM3 6.75C3 5.784 3.784 5 4.75 5h6.5c.966 0 1.75.784 1.75 1.75v1.5A1.75 1.75 0 0111.25 10h-6.5A1.75 1.75 0 013 8.25v-1.5zm1.47-.53a.75.75 0 011.06 0l.97.97.97-.97a.75.75 0 011.06 0l.97.97.97-.97a.75.75 0 111.06 1.06l-1.5 1.5a.75.75 0 01-1.06 0L8 7.81l-.97.97a.75.75 0 01-1.06 0l-1.5-1.5a.75.75 0 010-1.06z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-hubot%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M0 13C0 6.373 5.373 1 12 1s12 5.373 12 12v8.657a.75.75 0 01-1.5 0V13c0-5.799-4.701-10.5-10.5-10.5S1.5 7.201 1.5 13v8.657a.75.75 0 01-1.5 0V13z\"></path><path d=\"M8 19.75a.75.75 0 01.75-.75h6.5a.75.75 0 010 1.5h-6.5a.75.75 0 01-.75-.75z\"></path><path fill-rule=\"evenodd\" d=\"M5.25 9.5a1.75 1.75 0 00-1.75 1.75v3.5c0 .966.784 1.75 1.75 1.75h13.5a1.75 1.75 0 001.75-1.75v-3.5a1.75 1.75 0 00-1.75-1.75H5.25zm.22 1.47a.75.75 0 011.06 0L9 13.44l2.47-2.47a.75.75 0 011.06 0L15 13.44l2.47-2.47a.75.75 0 111.06 1.06l-3 3a.75.75 0 01-1.06 0L12 12.56l-2.47 2.47a.75.75 0 01-1.06 0l-3-3a.75.75 0 010-1.06z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// IdBadge returns a string representing an "id-badge" Octicon.
func IdBadge(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-id-badge%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M3 7.5a.5.5 0 01.5-.5h2a.5.5 0 01.5.5v3a.5.5 0 01-.5.5h-2a.5.5 0 01-.5-.5v-3zm10 .25a.75.75 0 01-.75.75h-4.5a.75.75 0 010-1.5h4.5a.75.75 0 01.75.75zM10.25 11a.75.75 0 000-1.5h-2.5a.75.75 0 000 1.5h2.5z\"></path><path fill-rule=\"evenodd\" d=\"M7.25 0A1.75 1.75 0 005.5 1.75V3H1.75A1.75 1.75 0 000 4.75v8.5C0 14.216.784 15 1.75 15h12.5A1.75 1.75 0 0016 13.25v-8.5A1.75 1.75 0 0014.25 3H10.5V1.75A1.75 1.75 0 008.75 0h-1.5zm3.232 4.5A1.75 1.75 0 018.75 6h-1.5a1.75 1.75 0 01-1.732-1.5H1.75a.25.25 0 00-.25.25v8.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25v-8.5a.25.25 0 00-.25-.25h-3.768zM7 1.75a.25.25 0 01.25-.25h1.5a.25.25 0 01.25.25v2.5a.25.25 0 01-.25.25h-1.5A.25.25 0 017 4.25v-2.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Image returns a string representing an "image" Octicon.
func Image(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-image%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 2.5a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h.94a.76.76 0 01.03-.03l6.077-6.078a1.75 1.75 0 012.412-.06L14.5 10.31V2.75a.25.25 0 00-.25-.25H1.75zm12.5 11H4.81l5.048-5.047a.25.25 0 01.344-.009l4.298 3.889v.917a.25.25 0 01-.25.25zm1.75-.25V2.75A1.75 1.75 0 0014.25 1H1.75A1.75 1.75 0 000 2.75v10.5C0 14.216.784 15 1.75 15h12.5A1.75 1.75 0 0016 13.25zM5.5 6a.5.5 0 11-1 0 .5.5 0 011 0zM7 6a2 2 0 11-4 0 2 2 0 014 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-image%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M19.25 4.5H4.75a.25.25 0 00-.25.25v14.5c0 .138.112.25.25.25h.19l9.823-9.823a1.75 1.75 0 012.475 0l2.262 2.262V4.75a.25.25 0 00-.25-.25zm.25 9.56l-3.323-3.323a.25.25 0 00-.354 0L7.061 19.5H19.25a.25.25 0 00.25-.25v-5.19zM4.75 3A1.75 1.75 0 003 4.75v14.5c0 .966.784 1.75 1.75 1.75h14.5A1.75 1.75 0 0021 19.25V4.75A1.75 1.75 0 0019.25 3H4.75zM8.5 9.5a1 1 0 100-2 1 1 0 000 2zm0 1.5a2.5 2.5 0 100-5 2.5 2.5 0 000 5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Inbox returns a string representing an "inbox" Octicon.
func Inbox(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-inbox%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.8 2.06A1.75 1.75 0 014.41 1h7.18c.7 0 1.333.417 1.61 1.06l2.74 6.395a.75.75 0 01.06.295v4.5A1.75 1.75 0 0114.25 15H1.75A1.75 1.75 0 010 13.25v-4.5a.75.75 0 01.06-.295L2.8 2.06zm1.61.44a.25.25 0 00-.23.152L1.887 8H4.75a.75.75 0 01.6.3L6.625 10h2.75l1.275-1.7a.75.75 0 01.6-.3h2.863L11.82 2.652a.25.25 0 00-.23-.152H4.41zm10.09 7h-2.875l-1.275 1.7a.75.75 0 01-.6.3h-3.5a.75.75 0 01-.6-.3L4.375 9.5H1.5v3.75c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V9.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-inbox%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.801 3.57A1.75 1.75 0 016.414 2.5h11.174c.702 0 1.337.42 1.611 1.067l3.741 8.828c.04.092.06.192.06.293v7.562A1.75 1.75 0 0121.25 22H2.75A1.75 1.75 0 011 20.25v-7.5c0-.1.02-.199.059-.291L4.8 3.571zM6.414 4a.25.25 0 00-.23.153L2.88 12H8a.75.75 0 01.648.372L10.18 15h3.638l1.533-2.628a.75.75 0 01.64-.372l5.13-.051-3.304-7.797a.25.25 0 00-.23-.152H6.414zM21.5 13.445l-5.067.05-1.535 2.633a.75.75 0 01-.648.372h-4.5a.75.75 0 01-.648-.372L7.57 13.5H2.5v6.75c0 .138.112.25.25.25h18.5a.25.25 0 00.25-.25v-6.805z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Infinity returns a string representing an "infinity" Octicon.
func Infinity(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-infinity%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.5 6c-1.086 0-2 .914-2 2 0 1.086.914 2 2 2 .525 0 1.122-.244 1.825-.727.51-.35 1.025-.79 1.561-1.273-.536-.483-1.052-.922-1.56-1.273C4.621 6.244 4.025 6 3.5 6zm4.5.984c-.59-.533-1.204-1.066-1.825-1.493-.797-.548-1.7-.991-2.675-.991C1.586 4.5 0 6.086 0 8s1.586 3.5 3.5 3.5c.975 0 1.878-.444 2.675-.991.621-.427 1.235-.96 1.825-1.493.59.533 1.204 1.066 1.825 1.493.797.547 1.7.991 2.675.991 1.914 0 3.5-1.586 3.5-3.5s-1.586-3.5-3.5-3.5c-.975 0-1.878.443-2.675.991-.621.427-1.235.96-1.825 1.493zM9.114 8c.536.483 1.052.922 1.56 1.273.704.483 1.3.727 1.826.727 1.086 0 2-.914 2-2 0-1.086-.914-2-2-2-.525 0-1.122.244-1.825.727-.51.35-1.025.79-1.561 1.273z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-infinity%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.25 8.5c-2.032 0-3.75 1.895-3.75 3.75S3.218 16 5.25 16c1.017 0 2.014-.457 3.062-1.253.89-.678 1.758-1.554 2.655-2.497-.897-.943-1.765-1.82-2.655-2.497C7.264 8.957 6.267 8.5 5.25 8.5zM12 11.16c-.887-.933-1.813-1.865-2.78-2.6C8.048 7.667 6.733 7 5.25 7 2.343 7 0 9.615 0 12.25s2.343 5.25 5.25 5.25c1.483 0 2.798-.668 3.97-1.56.967-.735 1.893-1.667 2.78-2.6.887.933 1.813 1.865 2.78 2.6 1.172.892 2.487 1.56 3.97 1.56 2.907 0 5.25-2.615 5.25-5.25S21.657 7 18.75 7c-1.483 0-2.798.668-3.97 1.56-.967.735-1.893 1.667-2.78 2.6zm1.033 1.09c.897.943 1.765 1.82 2.655 2.497C16.736 15.543 17.733 16 18.75 16c2.032 0 3.75-1.895 3.75-3.75S20.782 8.5 18.75 8.5c-1.017 0-2.014.457-3.062 1.253-.89.678-1.758 1.554-2.655 2.497z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Info returns a string representing an "info" Octicon.
func Info(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-info%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 1.5a6.5 6.5 0 100 13 6.5 6.5 0 000-13zM0 8a8 8 0 1116 0A8 8 0 010 8zm6.5-.25A.75.75 0 017.25 7h1a.75.75 0 01.75.75v2.75h.25a.75.75 0 010 1.5h-2a.75.75 0 010-1.5h.25v-2h-.25a.75.75 0 01-.75-.75zM8 6a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-info%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M13 7.5a1 1 0 11-2 0 1 1 0 012 0zm-3 3.75a.75.75 0 01.75-.75h1.5a.75.75 0 01.75.75v4.25h.75a.75.75 0 010 1.5h-3a.75.75 0 010-1.5h.75V12h-.75a.75.75 0 01-.75-.75z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// IssueClosed returns a string representing an "issue-closed" Octicon.
func IssueClosed(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-closed%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M11.28 6.78a.75.75 0 00-1.06-1.06L7.25 8.69 5.78 7.22a.75.75 0 00-1.06 1.06l2 2a.75.75 0 001.06 0l3.5-3.5z\"></path><path fill-rule=\"evenodd\" d=\"M16 8A8 8 0 110 8a8 8 0 0116 0zm-1.5 0a6.5 6.5 0 11-13 0 6.5 6.5 0 0113 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-closed%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M17.28 9.28a.75.75 0 00-1.06-1.06l-5.97 5.97-2.47-2.47a.75.75 0 00-1.06 1.06l3 3a.75.75 0 001.06 0l6.5-6.5z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// IssueDraft returns a string representing an "issue-draft" Octicon.
func IssueDraft(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-draft%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.749.097a8.054 8.054 0 012.502 0 .75.75 0 11-.233 1.482 6.554 6.554 0 00-2.036 0A.75.75 0 016.749.097zM4.345 1.693A.75.75 0 014.18 2.74a6.542 6.542 0 00-1.44 1.44.75.75 0 01-1.212-.883 8.042 8.042 0 011.769-1.77.75.75 0 011.048.166zm7.31 0a.75.75 0 011.048-.165 8.04 8.04 0 011.77 1.769.75.75 0 11-1.214.883 6.542 6.542 0 00-1.439-1.44.75.75 0 01-.165-1.047zM.955 6.125a.75.75 0 01.624.857 6.554 6.554 0 000 2.036.75.75 0 01-1.482.233 8.054 8.054 0 010-2.502.75.75 0 01.858-.624zm14.09 0a.75.75 0 01.858.624 8.057 8.057 0 010 2.502.75.75 0 01-1.482-.233 6.55 6.55 0 000-2.036.75.75 0 01.624-.857zm-13.352 5.53a.75.75 0 011.048.165 6.542 6.542 0 001.439 1.44.75.75 0 01-.883 1.212 8.04 8.04 0 01-1.77-1.769.75.75 0 01.166-1.048zm12.614 0a.75.75 0 01.165 1.048 8.038 8.038 0 01-1.769 1.77.75.75 0 11-.883-1.214 6.543 6.543 0 001.44-1.439.75.75 0 011.047-.165zm-8.182 3.39a.75.75 0 01.857-.624 6.55 6.55 0 002.036 0 .75.75 0 01.233 1.482 8.057 8.057 0 01-2.502 0 .75.75 0 01-.624-.858z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-draft%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.157 1.154a11.07 11.07 0 013.686 0 .75.75 0 01-.25 1.479 9.568 9.568 0 00-3.186 0 .75.75 0 01-.25-1.48zM6.68 3.205a.75.75 0 01-.177 1.046A9.558 9.558 0 004.25 6.503a.75.75 0 01-1.223-.87 11.058 11.058 0 012.606-2.605.75.75 0 011.046.177zm10.64 0a.75.75 0 011.046-.177 11.058 11.058 0 012.605 2.606.75.75 0 11-1.222.869 9.558 9.558 0 00-2.252-2.252.75.75 0 01-.177-1.046zM2.018 9.543a.75.75 0 01.615.864 9.568 9.568 0 000 3.186.75.75 0 01-1.48.25 11.07 11.07 0 010-3.686.75.75 0 01.865-.614zm19.964 0a.75.75 0 01.864.614 11.066 11.066 0 010 3.686.75.75 0 01-1.479-.25 9.56 9.56 0 000-3.186.75.75 0 01.615-.864zM3.205 17.32a.75.75 0 011.046.177 9.558 9.558 0 002.252 2.252.75.75 0 11-.87 1.223 11.058 11.058 0 01-2.605-2.606.75.75 0 01.177-1.046zm17.59 0a.75.75 0 01.176 1.046 11.057 11.057 0 01-2.605 2.605.75.75 0 11-.869-1.222 9.558 9.558 0 002.252-2.252.75.75 0 011.046-.177zM9.543 21.982a.75.75 0 01.864-.615 9.56 9.56 0 003.186 0 .75.75 0 01.25 1.48 11.066 11.066 0 01-3.686 0 .75.75 0 01-.614-.865z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// IssueOpened returns a string representing an "issue-opened" Octicon.
func IssueOpened(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-opened%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8 9.5a1.5 1.5 0 100-3 1.5 1.5 0 000 3z\"></path><path fill-rule=\"evenodd\" d=\"M8 0a8 8 0 100 16A8 8 0 008 0zM1.5 8a6.5 6.5 0 1113 0 6.5 6.5 0 01-13 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-opened%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0zM12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zm0 13a2 2 0 100-4 2 2 0 000 4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// IssueReopened returns a string representing an "issue-reopened" Octicon.
func IssueReopened(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-reopened%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M5.029 2.217a6.5 6.5 0 019.437 5.11.75.75 0 101.492-.154 8 8 0 00-14.315-4.03L.427 1.927A.25.25 0 000 2.104V5.75A.25.25 0 00.25 6h3.646a.25.25 0 00.177-.427L2.715 4.215a6.491 6.491 0 012.314-1.998zM1.262 8.169a.75.75 0 00-1.22.658 8.001 8.001 0 0014.315 4.03l1.216 1.216a.25.25 0 00.427-.177V10.25a.25.25 0 00-.25-.25h-3.646a.25.25 0 00-.177.427l1.358 1.358a6.501 6.501 0 01-11.751-3.11.75.75 0 00-.272-.506z\"></path><path d=\"M9.06 9.06a1.5 1.5 0 11-2.12-2.12 1.5 1.5 0 012.12 2.12z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-issue-reopened%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M3.38 8A9.502 9.502 0 0112 2.5a9.502 9.502 0 019.215 7.182.75.75 0 101.456-.364C21.473 4.539 17.15 1 12 1a10.995 10.995 0 00-9.5 5.452V4.75a.75.75 0 00-1.5 0V8.5a1 1 0 001 1h3.75a.75.75 0 000-1.5H3.38zm-.595 6.318a.75.75 0 00-1.455.364C2.527 19.461 6.85 23 12 23c4.052 0 7.592-2.191 9.5-5.451v1.701a.75.75 0 001.5 0V15.5a1 1 0 00-1-1h-3.75a.75.75 0 000 1.5h2.37A9.502 9.502 0 0112 21.5c-4.446 0-8.181-3.055-9.215-7.182z\"></path><path d=\"M13.414 13.414a2 2 0 11-2.828-2.828 2 2 0 012.828 2.828z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Italic returns a string representing an "italic" Octicon.
func Italic(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-italic%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 2.75A.75.75 0 016.75 2h6.5a.75.75 0 010 1.5h-2.505l-3.858 9H9.25a.75.75 0 010 1.5h-6.5a.75.75 0 010-1.5h2.505l3.858-9H6.75A.75.75 0 016 2.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-italic%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10 4.75a.75.75 0 01.75-.75h8.5a.75.75 0 010 1.5h-3.514l-5.828 13h3.342a.75.75 0 010 1.5h-8.5a.75.75 0 010-1.5h3.514l5.828-13H10.75a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Iterations returns a string representing an "iterations" Octicon.
func Iterations(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-iterations%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M2.5 7.25a4.75 4.75 0 019.5 0 .75.75 0 001.5 0 6.25 6.25 0 10-6.25 6.25H12v2.146c0 .223.27.335.427.177l2.896-2.896a.25.25 0 000-.354l-2.896-2.896a.25.25 0 00-.427.177V12H7.25A4.75 4.75 0 012.5 7.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-iterations%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M2.5 10.5a8 8 0 1116 0 .75.75 0 001.5 0 9.5 9.5 0 10-9.5 9.5h10.94l-2.72 2.72a.75.75 0 101.06 1.06l3.735-3.735c.44-.439.44-1.151 0-1.59L19.78 14.72a.75.75 0 00-1.06 1.06l2.72 2.72H10.5a8 8 0 01-8-8z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// KebabHorizontal returns a string representing an "kebab-horizontal" Octicon.
func KebabHorizontal(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-kebab-horizontal%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8 9a1.5 1.5 0 100-3 1.5 1.5 0 000 3zM1.5 9a1.5 1.5 0 100-3 1.5 1.5 0 000 3zm13 0a1.5 1.5 0 100-3 1.5 1.5 0 000 3z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-kebab-horizontal%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 12a2 2 0 11-4 0 2 2 0 014 0zm8 0a2 2 0 11-4 0 2 2 0 014 0zm6 2a2 2 0 100-4 2 2 0 000 4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Key returns a string representing an "key" Octicon.
func Key(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-key%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.5 5.5a4 4 0 112.731 3.795.75.75 0 00-.768.18L7.44 10.5H6.25a.75.75 0 00-.75.75v1.19l-.06.06H4.25a.75.75 0 00-.75.75v1.19l-.06.06H1.75a.25.25 0 01-.25-.25v-1.69l5.024-5.023a.75.75 0 00.181-.768A3.995 3.995 0 016.5 5.5zm4-5.5a5.5 5.5 0 00-5.348 6.788L.22 11.72a.75.75 0 00-.22.53v2C0 15.216.784 16 1.75 16h2a.75.75 0 00.53-.22l.5-.5a.75.75 0 00.22-.53V14h.75a.75.75 0 00.53-.22l.5-.5a.75.75 0 00.22-.53V12h.75a.75.75 0 00.53-.22l.932-.932A5.5 5.5 0 1010.5 0zm.5 6a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-key%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M16.75 8.5a1.25 1.25 0 100-2.5 1.25 1.25 0 000 2.5z\"></path><path fill-rule=\"evenodd\" d=\"M15.75 0a8.25 8.25 0 00-7.851 10.79L.513 18.178A1.75 1.75 0 000 19.414v2.836C0 23.217.784 24 1.75 24h1.5A1.75 1.75 0 005 22.25v-1a.25.25 0 01.25-.25h2.735a.75.75 0 00.545-.22l.214-.213A.875.875 0 009 19.948V18.5a.25.25 0 01.25-.25h1.086c.464 0 .91-.184 1.237-.513l1.636-1.636A8.25 8.25 0 1015.75 0zM9 8.25a6.75 6.75 0 114.288 6.287.75.75 0 00-.804.168l-1.971 1.972a.25.25 0 01-.177.073H9.25A1.75 1.75 0 007.5 18.5v1H5.25a1.75 1.75 0 00-1.75 1.75v1a.25.25 0 01-.25.25h-1.5a.25.25 0 01-.25-.25v-2.836a.25.25 0 01.073-.177l7.722-7.721a.75.75 0 00.168-.804A6.73 6.73 0 019 8.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// KeyAsterisk returns a string representing an "key-asterisk" Octicon.
func KeyAsterisk(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-key-asterisk%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 2.75A2.75 2.75 0 012.75 0h10.5A2.75 2.75 0 0116 2.75v10.5A2.75 2.75 0 0113.25 16H2.75A2.75 2.75 0 010 13.25V2.75zM2.75 1.5c-.69 0-1.25.56-1.25 1.25v10.5c0 .69.56 1.25 1.25 1.25h10.5c.69 0 1.25-.56 1.25-1.25V2.75c0-.69-.56-1.25-1.25-1.25H2.75z\"></path><path d=\"M8 4a.75.75 0 01.75.75V6.7l1.69-.975a.75.75 0 01.75 1.3L9.5 8l1.69.976a.75.75 0 01-.75 1.298L8.75 9.3v1.951a.75.75 0 01-1.5 0V9.299l-1.69.976a.75.75 0 01-.75-1.3L6.5 8l-1.69-.975a.75.75 0 01.75-1.3l1.69.976V4.75A.75.75 0 018 4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Law returns a string representing an "law" Octicon.
func Law(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-law%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.75.75a.75.75 0 00-1.5 0V2h-.984c-.305 0-.604.08-.869.23l-1.288.737A.25.25 0 013.984 3H1.75a.75.75 0 000 1.5h.428L.066 9.192a.75.75 0 00.154.838l.53-.53-.53.53v.001l.002.002.002.002.006.006.016.015.045.04a3.514 3.514 0 00.686.45A4.492 4.492 0 003 11c.88 0 1.556-.22 2.023-.454a3.515 3.515 0 00.686-.45l.045-.04.016-.015.006-.006.002-.002.001-.002L5.25 9.5l.53.53a.75.75 0 00.154-.838L3.822 4.5h.162c.305 0 .604-.08.869-.23l1.289-.737a.25.25 0 01.124-.033h.984V13h-2.5a.75.75 0 000 1.5h6.5a.75.75 0 000-1.5h-2.5V3.5h.984a.25.25 0 01.124.033l1.29.736c.264.152.563.231.868.231h.162l-2.112 4.692a.75.75 0 00.154.838l.53-.53-.53.53v.001l.002.002.002.002.006.006.016.015.045.04a3.517 3.517 0 00.686.45A4.492 4.492 0 0013 11c.88 0 1.556-.22 2.023-.454a3.512 3.512 0 00.686-.45l.045-.04.01-.01.006-.005.006-.006.002-.002.001-.002-.529-.531.53.53a.75.75 0 00.154-.838L13.823 4.5h.427a.75.75 0 000-1.5h-2.234a.25.25 0 01-.124-.033l-1.29-.736A1.75 1.75 0 009.735 2H8.75V.75zM1.695 9.227c.285.135.718.273 1.305.273s1.02-.138 1.305-.273L3 6.327l-1.305 2.9zm10 0c.285.135.718.273 1.305.273s1.02-.138 1.305-.273L13 6.327l-1.305 2.9z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-law%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.75 2.75a.75.75 0 00-1.5 0V4.5H9.276a1.75 1.75 0 00-.985.303L6.596 5.957A.25.25 0 016.455 6H2.353a.75.75 0 100 1.5H3.93L.563 15.18a.762.762 0 00.21.88c.08.064.161.125.309.221.186.121.452.278.792.433.68.311 1.662.62 2.876.62a6.919 6.919 0 002.876-.62c.34-.155.606-.312.792-.433.15-.097.23-.158.31-.223a.75.75 0 00.209-.878L5.569 7.5h.886c.351 0 .694-.106.984-.303l1.696-1.154A.25.25 0 019.275 6h1.975v14.5H6.763a.75.75 0 000 1.5h10.474a.75.75 0 000-1.5H12.75V6h1.974c.05 0 .1.015.14.043l1.697 1.154c.29.197.633.303.984.303h.886l-3.368 7.68a.75.75 0 00.23.896c.012.009 0 0 .002 0a3.154 3.154 0 00.31.206c.185.112.45.256.79.4a7.343 7.343 0 002.855.568 7.343 7.343 0 002.856-.569c.338-.143.604-.287.79-.399a3.5 3.5 0 00.31-.206.75.75 0 00.23-.896L20.07 7.5h1.578a.75.75 0 000-1.5h-4.102a.25.25 0 01-.14-.043l-1.697-1.154a1.75 1.75 0 00-.984-.303H12.75V2.75zM2.193 15.198a5.418 5.418 0 002.557.635 5.418 5.418 0 002.557-.635L4.75 9.368l-2.557 5.83zm14.51-.024c.082.04.174.083.275.126.53.223 1.305.45 2.272.45a5.846 5.846 0 002.547-.576L19.25 9.367l-2.547 5.807z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// LightBulb returns a string representing an "light-bulb" Octicon.
func LightBulb(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-light-bulb%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 1.5c-2.363 0-4 1.69-4 3.75 0 .984.424 1.625.984 2.304l.214.253c.223.264.47.556.673.848.284.411.537.896.621 1.49a.75.75 0 01-1.484.211c-.04-.282-.163-.547-.37-.847a8.695 8.695 0 00-.542-.68c-.084-.1-.173-.205-.268-.32C3.201 7.75 2.5 6.766 2.5 5.25 2.5 2.31 4.863 0 8 0s5.5 2.31 5.5 5.25c0 1.516-.701 2.5-1.328 3.259-.095.115-.184.22-.268.319-.207.245-.383.453-.541.681-.208.3-.33.565-.37.847a.75.75 0 01-1.485-.212c.084-.593.337-1.078.621-1.489.203-.292.45-.584.673-.848.075-.088.147-.173.213-.253.561-.679.985-1.32.985-2.304 0-2.06-1.637-3.75-4-3.75zM6 15.25a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5h-2.5a.75.75 0 01-.75-.75zM5.75 12a.75.75 0 000 1.5h4.5a.75.75 0 000-1.5h-4.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-light-bulb%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 2.5c-3.81 0-6.5 2.743-6.5 6.119 0 1.536.632 2.572 1.425 3.56.172.215.347.422.527.635l.096.112c.21.25.427.508.63.774.404.531.783 1.128.995 1.834a.75.75 0 01-1.436.432c-.138-.46-.397-.89-.753-1.357a18.354 18.354 0 00-.582-.714l-.092-.11c-.18-.212-.37-.436-.555-.667C4.87 12.016 4 10.651 4 8.618 4 4.363 7.415 1 12 1s8 3.362 8 7.619c0 2.032-.87 3.397-1.755 4.5-.185.23-.375.454-.555.667l-.092.109c-.21.248-.405.481-.582.714-.356.467-.615.898-.753 1.357a.75.75 0 01-1.437-.432c.213-.706.592-1.303.997-1.834.202-.266.419-.524.63-.774l.095-.112c.18-.213.355-.42.527-.634.793-.99 1.425-2.025 1.425-3.561C18.5 5.243 15.81 2.5 12 2.5zM9.5 21.75a.75.75 0 01.75-.75h3.5a.75.75 0 010 1.5h-3.5a.75.75 0 01-.75-.75zM8.75 18a.75.75 0 000 1.5h6.5a.75.75 0 000-1.5h-6.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Link returns a string representing an "link" Octicon.
func Link(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-link%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.775 3.275a.75.75 0 001.06 1.06l1.25-1.25a2 2 0 112.83 2.83l-2.5 2.5a2 2 0 01-2.83 0 .75.75 0 00-1.06 1.06 3.5 3.5 0 004.95 0l2.5-2.5a3.5 3.5 0 00-4.95-4.95l-1.25 1.25zm-4.69 9.64a2 2 0 010-2.83l2.5-2.5a2 2 0 012.83 0 .75.75 0 001.06-1.06 3.5 3.5 0 00-4.95 0l-2.5 2.5a3.5 3.5 0 004.95 4.95l1.25-1.25a.75.75 0 00-1.06-1.06l-1.25 1.25a2 2 0 01-2.83 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-link%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M14.78 3.653a3.936 3.936 0 115.567 5.567l-3.627 3.627a3.936 3.936 0 01-5.88-.353.75.75 0 00-1.18.928 5.436 5.436 0 008.12.486l3.628-3.628a5.436 5.436 0 10-7.688-7.688l-3 3a.75.75 0 001.06 1.061l3-3z\"></path><path d=\"M7.28 11.153a3.936 3.936 0 015.88.353.75.75 0 001.18-.928 5.436 5.436 0 00-8.12-.486L2.592 13.72a5.436 5.436 0 107.688 7.688l3-3a.75.75 0 10-1.06-1.06l-3 3a3.936 3.936 0 01-5.567-5.568l3.627-3.627z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// LinkExternal returns a string representing an "link-external" Octicon.
func LinkExternal(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-link-external%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.604 1h4.146a.25.25 0 01.25.25v4.146a.25.25 0 01-.427.177L13.03 4.03 9.28 7.78a.75.75 0 01-1.06-1.06l3.75-3.75-1.543-1.543A.25.25 0 0110.604 1zM3.75 2A1.75 1.75 0 002 3.75v8.5c0 .966.784 1.75 1.75 1.75h8.5A1.75 1.75 0 0014 12.25v-3.5a.75.75 0 00-1.5 0v3.5a.25.25 0 01-.25.25h-8.5a.25.25 0 01-.25-.25v-8.5a.25.25 0 01.25-.25h3.5a.75.75 0 000-1.5h-3.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-link-external%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M15.5 2.25a.75.75 0 01.75-.75h5.5a.75.75 0 01.75.75v5.5a.75.75 0 01-1.5 0V4.06l-6.22 6.22a.75.75 0 11-1.06-1.06L19.94 3h-3.69a.75.75 0 01-.75-.75z\"></path><path d=\"M2.5 4.25c0-.966.784-1.75 1.75-1.75h8.5a.75.75 0 010 1.5h-8.5a.25.25 0 00-.25.25v15.5c0 .138.112.25.25.25h15.5a.25.25 0 00.25-.25v-8.5a.75.75 0 011.5 0v8.5a1.75 1.75 0 01-1.75 1.75H4.25a1.75 1.75 0 01-1.75-1.75V4.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ListOrdered returns a string representing an "list-ordered" Octicon.
func ListOrdered(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-list-ordered%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.003 2.5a.5.5 0 00-.723-.447l-1.003.5a.5.5 0 00.446.895l.28-.14V6H.5a.5.5 0 000 1h2.006a.5.5 0 100-1h-.503V2.5zM5 3.25a.75.75 0 01.75-.75h8.5a.75.75 0 010 1.5h-8.5A.75.75 0 015 3.25zm0 5a.75.75 0 01.75-.75h8.5a.75.75 0 010 1.5h-8.5A.75.75 0 015 8.25zm0 5a.75.75 0 01.75-.75h8.5a.75.75 0 010 1.5h-8.5a.75.75 0 01-.75-.75zM.924 10.32l.003-.004a.851.851 0 01.144-.153A.66.66 0 011.5 10c.195 0 .306.068.374.146a.57.57 0 01.128.376c0 .453-.269.682-.8 1.078l-.035.025C.692 11.98 0 12.495 0 13.5a.5.5 0 00.5.5h2.003a.5.5 0 000-1H1.146c.132-.197.351-.372.654-.597l.047-.035c.47-.35 1.156-.858 1.156-1.845 0-.365-.118-.744-.377-1.038-.268-.303-.658-.484-1.126-.484-.48 0-.84.202-1.068.392a1.858 1.858 0 00-.348.384l-.007.011-.002.004-.001.002-.001.001a.5.5 0 00.851.525zM.5 10.055l-.427-.26.427.26z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-list-ordered%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M3.604 3.089A.75.75 0 014 3.75V8.5h.75a.75.75 0 010 1.5h-3a.75.75 0 110-1.5h.75V5.151l-.334.223a.75.75 0 01-.832-1.248l1.5-1a.75.75 0 01.77-.037zM8.75 5.5a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zm0 6a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zm0 6a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zM5.5 15.75c0-.704-.271-1.286-.72-1.686a2.302 2.302 0 00-1.53-.564c-.535 0-1.094.178-1.53.565-.449.399-.72.982-.72 1.685a.75.75 0 001.5 0c0-.296.104-.464.217-.564A.805.805 0 013.25 15c.215 0 .406.072.533.185.113.101.217.268.217.565 0 .332-.069.48-.21.657-.092.113-.216.24-.403.419l-.147.14c-.152.144-.33.313-.52.504l-1.5 1.5a.75.75 0 00-.22.53v.25c0 .414.336.75.75.75H5A.75.75 0 005 19H3.31l.47-.47c.176-.176.333-.324.48-.465l.165-.156a5.98 5.98 0 00.536-.566c.358-.447.539-.925.539-1.593z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ListUnordered returns a string representing an "list-unordered" Octicon.
func ListUnordered(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-list-unordered%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 4a1 1 0 100-2 1 1 0 000 2zm3.75-1.5a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5zm0 5a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5zm0 5a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5zM3 8a1 1 0 11-2 0 1 1 0 012 0zm-1 6a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-list-unordered%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4 7a1 1 0 100-2 1 1 0 000 2zm4.75-1.5a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zm0 6a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zm0 6a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zM5 12a1 1 0 11-2 0 1 1 0 012 0zm-1 7a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Location returns a string representing an "location" Octicon.
func Location(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-location%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.536 3.464a5 5 0 010 7.072L8 14.07l-3.536-3.535a5 5 0 117.072-7.072v.001zm1.06 8.132a6.5 6.5 0 10-9.192 0l3.535 3.536a1.5 1.5 0 002.122 0l3.535-3.536zM8 9a2 2 0 100-4 2 2 0 000 4z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-location%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12 13.5a2.5 2.5 0 100-5 2.5 2.5 0 000 5z\"></path><path fill-rule=\"evenodd\" d=\"M19.071 3.429C15.166-.476 8.834-.476 4.93 3.429c-3.905 3.905-3.905 10.237 0 14.142l.028.028 5.375 5.375a2.359 2.359 0 003.336 0l5.403-5.403c3.905-3.905 3.905-10.237 0-14.142zM5.99 4.489A8.5 8.5 0 0118.01 16.51l-5.403 5.404a.859.859 0 01-1.214 0l-5.378-5.378-.002-.002-.023-.024a8.5 8.5 0 010-12.02z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Lock returns a string representing an "lock" Octicon.
func Lock(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-lock%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4 4v2h-.25A1.75 1.75 0 002 7.75v5.5c0 .966.784 1.75 1.75 1.75h8.5A1.75 1.75 0 0014 13.25v-5.5A1.75 1.75 0 0012.25 6H12V4a4 4 0 10-8 0zm6.5 2V4a2.5 2.5 0 00-5 0v2h5zM12 7.5h.25a.25.25 0 01.25.25v5.5a.25.25 0 01-.25.25h-8.5a.25.25 0 01-.25-.25v-5.5a.25.25 0 01.25-.25H12z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-lock%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 9V7.25C6 3.845 8.503 1 12 1s6 2.845 6 6.25V9h.5a2.5 2.5 0 012.5 2.5v8a2.5 2.5 0 01-2.5 2.5h-13A2.5 2.5 0 013 19.5v-8A2.5 2.5 0 015.5 9H6zm1.5-1.75C7.5 4.58 9.422 2.5 12 2.5c2.578 0 4.5 2.08 4.5 4.75V9h-9V7.25zm-3 4.25a1 1 0 011-1h13a1 1 0 011 1v8a1 1 0 01-1 1h-13a1 1 0 01-1-1v-8z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Log returns a string representing an "log" Octicon.
func Log(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-log%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M5 8.25a.75.75 0 01.75-.75h4a.75.75 0 010 1.5h-4A.75.75 0 015 8.25zM4 10.5A.75.75 0 004 12h4a.75.75 0 000-1.5H4z\"></path><path fill-rule=\"evenodd\" d=\"M13-.005H3a3 3 0 00-3 3c0 .676.224 1.254.603 1.722.526.65 1.331.783 1.907.783h1.177c-.364.662-.814 1.339-1.287 2.048-.205.309-.414.624-.623.946C.891 9.865 0 11.418 0 13a3 3 0 003 3h10a3 3 0 001.667-5.494.75.75 0 00-.834 1.246A1.5 1.5 0 1111.5 13c0-.642.225-1.347.623-2.136.397-.787.933-1.593 1.501-2.446l.011-.017c.554-.83 1.139-1.709 1.582-2.588.445-.885.783-1.836.783-2.818 0-1.672-1.346-3-3-3zm-10 1.5a1.5 1.5 0 00-1.5 1.5c0 .321.1.569.27.778.097.12.325.227.74.227h7.674A2.737 2.737 0 0110 2.995c0-.546.146-1.059.401-1.5H3zm10 0c.831 0 1.5.662 1.5 1.5 0 .646-.225 1.353-.623 2.143-.398.79-.933 1.595-1.501 2.448l-.017.026c-.552.828-1.134 1.702-1.575 2.576C10.338 11.072 10 12.021 10 13c0 .546.146 1.059.401 1.5H3A1.5 1.5 0 011.5 13c0-1.084.63-2.289 1.537-3.692.177-.274.366-.556.558-.845.632-.948 1.306-1.96 1.773-2.963h6.382a.75.75 0 00.417-1.373c-.444-.298-.667-.656-.667-1.132a1.5 1.5 0 011.5-1.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// LogoGist returns a string representing an "logo-gist" Octicon.
func LogoGist(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-logo-gist%s\" height=\"16\" width=\"25\" viewbox=\"0 0 16 25\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.7 8.73h2.45v4.02c-.55.27-1.64.34-2.53.34-2.56 0-3.47-2.2-3.47-5.05 0-2.85.91-5.06 3.48-5.06 1.28 0 2.06.23 3.28.73V2.66C7.27 2.33 6.25 2 4.63 2 1.13 2 0 4.69 0 8.03c0 3.34 1.11 6.03 4.63 6.03 1.64 0 2.81-.27 3.59-.64V7.73H4.7v1zm6.39 3.72V6.06h-1.05v6.28c0 1.25.58 1.72 1.72 1.72v-.89c-.48 0-.67-.16-.67-.7v-.02zm.25-8.72c0-.44-.33-.78-.78-.78s-.77.34-.77.78.33.78.77.78.78-.34.78-.78zm4.34 5.69c-1.5-.13-1.78-.48-1.78-1.17 0-.77.33-1.34 1.88-1.34 1.05 0 1.66.16 2.27.36v-.94c-.69-.3-1.52-.39-2.25-.39-2.2 0-2.92 1.2-2.92 2.31 0 1.08.47 1.88 2.73 2.08 1.55.13 1.77.63 1.77 1.34 0 .73-.44 1.42-2.06 1.42-1.11 0-1.86-.19-2.33-.36v.94c.5.2 1.58.39 2.33.39 2.38 0 3.14-1.2 3.14-2.41 0-1.28-.53-2.03-2.75-2.23h-.03zm8.58-2.47v-.86h-2.42v-2.5l-1.08.31v2.11l-1.56.44v.48h1.56v5c0 1.53 1.19 2.13 2.5 2.13.19 0 .52-.02.69-.05v-.89c-.19.03-.41.03-.61.03-.97 0-1.5-.39-1.5-1.34V6.94h2.42v.02-.01z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// LogoGithub returns a string representing an "logo-github" Octicon.
func LogoGithub(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-logo-github%s\" height=\"16\" width=\"45\" viewbox=\"0 0 16 45\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M18.53 12.03h-.02c.009 0 .015.01.024.011h.006l-.01-.01zm.004.011c-.093.001-.327.05-.574.05-.78 0-1.05-.36-1.05-.83V8.13h1.59c.09 0 .16-.08.16-.19v-1.7c0-.09-.08-.17-.16-.17h-1.59V3.96c0-.08-.05-.13-.14-.13h-2.16c-.09 0-.14.05-.14.13v2.17s-1.09.27-1.16.28c-.08.02-.13.09-.13.17v1.36c0 .11.08.19.17.19h1.11v3.28c0 2.44 1.7 2.69 2.86 2.69.53 0 1.17-.17 1.27-.22.06-.02.09-.09.09-.16v-1.5a.177.177 0 00-.146-.18zM42.23 9.84c0-1.81-.73-2.05-1.5-1.97-.6.04-1.08.34-1.08.34v3.52s.49.34 1.22.36c1.03.03 1.36-.34 1.36-2.25zm2.43-.16c0 3.43-1.11 4.41-3.05 4.41-1.64 0-2.52-.83-2.52-.83s-.04.46-.09.52c-.03.06-.08.08-.14.08h-1.48c-.1 0-.19-.08-.19-.17l.02-11.11c0-.09.08-.17.17-.17h2.13c.09 0 .17.08.17.17v3.77s.82-.53 2.02-.53l-.01-.02c1.2 0 2.97.45 2.97 3.88zm-8.72-3.61h-2.1c-.11 0-.17.08-.17.19v5.44s-.55.39-1.3.39-.97-.34-.97-1.09V6.25c0-.09-.08-.17-.17-.17h-2.14c-.09 0-.17.08-.17.17v5.11c0 2.2 1.23 2.75 2.92 2.75 1.39 0 2.52-.77 2.52-.77s.05.39.08.45c.02.05.09.09.16.09h1.34c.11 0 .17-.08.17-.17l.02-7.47c0-.09-.08-.17-.19-.17zm-23.7-.01h-2.13c-.09 0-.17.09-.17.2v7.34c0 .2.13.27.3.27h1.92c.2 0 .25-.09.25-.27V6.23c0-.09-.08-.17-.17-.17zm-1.05-3.38c-.77 0-1.38.61-1.38 1.38 0 .77.61 1.38 1.38 1.38.75 0 1.36-.61 1.36-1.38 0-.77-.61-1.38-1.36-1.38zm16.49-.25h-2.11c-.09 0-.17.08-.17.17v4.09h-3.31V2.6c0-.09-.08-.17-.17-.17h-2.13c-.09 0-.17.08-.17.17v11.11c0 .09.09.17.17.17h2.13c.09 0 .17-.08.17-.17V8.96h3.31l-.02 4.75c0 .09.08.17.17.17h2.13c.09 0 .17-.08.17-.17V2.6c0-.09-.08-.17-.17-.17zM8.81 7.35v5.74c0 .04-.01.11-.06.13 0 0-1.25.89-3.31.89-2.49 0-5.44-.78-5.44-5.92S2.58 1.99 5.1 2c2.18 0 3.06.49 3.2.58.04.05.06.09.06.14L7.94 4.5c0 .09-.09.2-.2.17-.36-.11-.9-.33-2.17-.33-1.47 0-3.05.42-3.05 3.73s1.5 3.7 2.58 3.7c.92 0 1.25-.11 1.25-.11v-2.3H4.88c-.11 0-.19-.08-.19-.17V7.35c0-.09.08-.17.19-.17h3.74c.11 0 .19.08.19.17z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Mail returns a string representing an "mail" Octicon.
func Mail(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-mail%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 2A1.75 1.75 0 000 3.75v.736a.75.75 0 000 .027v7.737C0 13.216.784 14 1.75 14h12.5A1.75 1.75 0 0016 12.25v-8.5A1.75 1.75 0 0014.25 2H1.75zM14.5 4.07v-.32a.25.25 0 00-.25-.25H1.75a.25.25 0 00-.25.25v.32L8 7.88l6.5-3.81zm-13 1.74v6.441c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V5.809L8.38 9.397a.75.75 0 01-.76 0L1.5 5.809z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-mail%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 3A1.75 1.75 0 000 4.75v14c0 .966.784 1.75 1.75 1.75h20.5A1.75 1.75 0 0024 18.75v-14A1.75 1.75 0 0022.25 3H1.75zM1.5 4.75a.25.25 0 01.25-.25h20.5a.25.25 0 01.25.25v.852l-10.36 7a.25.25 0 01-.28 0l-10.36-7V4.75zm0 2.662V18.75c0 .138.112.25.25.25h20.5a.25.25 0 00.25-.25V7.412l-9.52 6.433c-.592.4-1.368.4-1.96 0L1.5 7.412z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// MarkGithub returns a string representing an "mark-github" Octicon.
func MarkGithub(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-mark-github%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Markdown returns a string representing an "markdown" Octicon.
func Markdown(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-markdown%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M14.85 3H1.15C.52 3 0 3.52 0 4.15v7.69C0 12.48.52 13 1.15 13h13.69c.64 0 1.15-.52 1.15-1.15v-7.7C16 3.52 15.48 3 14.85 3zM9 11H7V8L5.5 9.92 4 8v3H2V5h2l1.5 2L7 5h2v6zm2.99.5L9.5 8H11V5h2v3h1.5l-2.51 3.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Megaphone returns a string representing an "megaphone" Octicon.
func Megaphone(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-megaphone%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><g fill-rule=\"evenodd\"><path d=\"M3.25 9a.75.75 0 01.75.75c0 2.142.456 3.828.733 4.653a.121.121 0 00.05.064.207.207 0 00.117.033h1.31c.085 0 .18-.042.258-.152a.448.448 0 00.075-.366A16.74 16.74 0 016 9.75a.75.75 0 011.5 0c0 1.588.25 2.926.494 3.85.293 1.113-.504 2.4-1.783 2.4H4.9c-.686 0-1.35-.41-1.589-1.12A16.42 16.42 0 012.5 9.75.75.75 0 013.25 9z\"></path><path d=\"M0 6a4 4 0 014-4h2.75a.75.75 0 01.75.75v6.5a.75.75 0 01-.75.75H4a4 4 0 01-4-4zm4-2.5a2.5 2.5 0 000 5h2v-5H4z\"></path><path d=\"M15.59.082A.75.75 0 0116 .75v10.5a.75.75 0 01-1.189.608l-.002-.001h.001l-.014-.01a5.829 5.829 0 00-.422-.25 10.58 10.58 0 00-1.469-.64C11.576 10.484 9.536 10 6.75 10a.75.75 0 110-1.5c2.964 0 5.174.516 6.658 1.043.423.151.787.302 1.092.443V2.014c-.305.14-.669.292-1.092.443C11.924 2.984 9.713 3.5 6.75 3.5a.75.75 0 110-1.5c2.786 0 4.826-.484 6.155-.957.665-.236 1.154-.47 1.47-.64a5.82 5.82 0 00.421-.25l.014-.01a.75.75 0 01.78-.061zm-.78.06zm.44 11.108l-.44.607.44-.607z\"></path></g></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-megaphone%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M22 1.75a.75.75 0 00-1.161-.627c-.047.03-.094.057-.142.085a9.15 9.15 0 01-.49.262c-.441.22-1.11.519-2.002.82-1.78.6-4.45 1.21-7.955 1.21H6.5A5.5 5.5 0 005 14.293v.457c0 3.061.684 5.505 1.061 6.621.24.709.904 1.129 1.6 1.129h2.013c1.294 0 2.1-1.322 1.732-2.453-.412-1.268-.906-3.268-.906-5.547 0-.03-.002-.06-.005-.088 3.382.028 5.965.644 7.703 1.251.89.312 1.559.62 2 .849.084.043.171.096.261.15.357.214.757.455 1.142.25A.75.75 0 0022 16.25V1.75zM10.5 12.912c3.564.029 6.313.678 8.193 1.335.737.258 1.34.517 1.807.74V2.993c-.467.216-1.073.467-1.815.718-1.878.634-4.624 1.26-8.185 1.288v7.913zm-4 1.838v-.25H9c0 2.486.537 4.648.98 6.01a.398.398 0 01-.057.343c-.07.104-.162.147-.249.147H7.661c-.105 0-.161-.058-.179-.109-.344-1.018-.982-3.294-.982-6.141zM6.5 5H9v8H6.5a4 4 0 010-8z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Mention returns a string representing an "mention" Octicon.
func Mention(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-mention%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 2.37a6.5 6.5 0 006.5 11.26.75.75 0 01.75 1.298 8 8 0 113.994-7.273.754.754 0 01.006.095v1.5a2.75 2.75 0 01-5.072 1.475A4 4 0 1112 8v1.25a1.25 1.25 0 002.5 0V7.867a6.5 6.5 0 00-9.75-5.496V2.37zM10.5 8a2.5 2.5 0 10-5 0 2.5 2.5 0 005 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-mention%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M20.226 7.25a9.498 9.498 0 10-3.477 12.976.75.75 0 01.75 1.299c-5.26 3.037-11.987 1.235-15.024-4.026C-.562 12.24 1.24 5.512 6.501 2.475 11.76-.562 18.488 1.24 21.525 6.501a10.956 10.956 0 011.455 4.826c.013.056.02.113.02.173v2.25a3.5 3.5 0 01-6.623 1.581 5.5 5.5 0 111.112-3.682.76.76 0 01.011.129v1.972a2 2 0 104 0v-1.766a9.452 9.452 0 00-1.274-4.733zM16 12a4 4 0 10-8 0 4 4 0 008 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Meter returns a string representing an "meter" Octicon.
func Meter(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-meter%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 1.5a6.5 6.5 0 106.016 4.035.75.75 0 011.388-.57 8 8 0 11-4.37-4.37.75.75 0 01-.569 1.389A6.479 6.479 0 008 1.5zm6.28.22a.75.75 0 010 1.06l-4.063 4.064a2.5 2.5 0 11-1.06-1.06L13.22 1.72a.75.75 0 011.06 0zM7 8a1 1 0 112 0 1 1 0 01-2 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Milestone returns a string representing an "milestone" Octicon.
func Milestone(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-milestone%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.75 0a.75.75 0 01.75.75V3h3.634c.414 0 .814.147 1.13.414l2.07 1.75a1.75 1.75 0 010 2.672l-2.07 1.75a1.75 1.75 0 01-1.13.414H8.5v5.25a.75.75 0 11-1.5 0V10H2.75A1.75 1.75 0 011 8.25v-3.5C1 3.784 1.784 3 2.75 3H7V.75A.75.75 0 017.75 0zm0 8.5h4.384a.25.25 0 00.161-.06l2.07-1.75a.25.25 0 000-.38l-2.07-1.75a.25.25 0 00-.161-.06H2.75a.25.25 0 00-.25.25v3.5c0 .138.112.25.25.25h5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-milestone%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.75 1a.75.75 0 01.75.75V4h6.532c.42 0 .826.15 1.143.425l3.187 2.75a1.75 1.75 0 010 2.65l-3.187 2.75a1.75 1.75 0 01-1.143.425H12.5v9.25a.75.75 0 01-1.5 0V13H3.75A1.75 1.75 0 012 11.25v-5.5C2 4.783 2.784 4 3.75 4H11V1.75a.75.75 0 01.75-.75zm0 4.5h7.282a.25.25 0 01.163.06l3.188 2.75a.25.25 0 010 .38l-3.188 2.75a.25.25 0 01-.163.06H3.75a.25.25 0 01-.25-.25v-5.5a.25.25 0 01.25-.25h8z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Mirror returns a string representing an "mirror" Octicon.
func Mirror(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-mirror%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.75 1.75a.75.75 0 00-1.5 0v.5a.75.75 0 001.5 0v-.5zM8 4a.75.75 0 01.75.75v.5a.75.75 0 01-1.5 0v-.5A.75.75 0 018 4zm.75 3.75a.75.75 0 00-1.5 0v.5a.75.75 0 001.5 0v-.5zM8 10a.75.75 0 01.75.75v.5a.75.75 0 01-1.5 0v-.5A.75.75 0 018 10zm0 3a.75.75 0 01.75.75v.5a.75.75 0 01-1.5 0v-.5A.75.75 0 018 13zm7.547-9.939A.75.75 0 0116 3.75v8.5a.75.75 0 01-1.265.545l-4.5-4.25a.75.75 0 010-1.09l4.5-4.25a.75.75 0 01.812-.144zM11.842 8l2.658 2.51V5.49L11.842 8zM0 12.25a.75.75 0 001.265.545l4.5-4.25a.75.75 0 000-1.09l-4.5-4.25A.75.75 0 000 3.75v8.5zm1.5-6.76L4.158 8 1.5 10.51V5.49z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-mirror%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 10.75a.75.75 0 01.75.75v1a.75.75 0 01-1.5 0v-1a.75.75 0 01.75-.75zm0 4a.75.75 0 01.75.75v1a.75.75 0 01-1.5 0v-1a.75.75 0 01.75-.75zm0 4a.75.75 0 01.75.75v1a.75.75 0 01-1.5 0v-1a.75.75 0 01.75-.75zm0-12a.75.75 0 01.75.75v1a.75.75 0 01-1.5 0v-1a.75.75 0 01.75-.75zm0-4a.75.75 0 01.75.75v1a.75.75 0 01-1.5 0v-1a.75.75 0 01.75-.75zm9.553 3.314A.75.75 0 0122 6.75v10.5a.75.75 0 01-1.256.554l-5.75-5.25a.75.75 0 010-1.108l5.75-5.25a.75.75 0 01.809-.132zM16.613 12l3.887 3.55v-7.1L16.612 12zM2.447 17.936A.75.75 0 012 17.25V6.75a.75.75 0 011.256-.554l5.75 5.25a.75.75 0 010 1.108l-5.75 5.25a.75.75 0 01-.809.132zM7.387 12L3.5 8.45v7.1L7.388 12z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Moon returns a string representing an "moon" Octicon.
func Moon(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-moon%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.598 1.591a.75.75 0 01.785-.175 7 7 0 11-8.967 8.967.75.75 0 01.961-.96 5.5 5.5 0 007.046-7.046.75.75 0 01.175-.786zm1.616 1.945a7 7 0 01-7.678 7.678 5.5 5.5 0 107.678-7.678z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-moon%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M16.5 6c0 5.799-4.701 10.5-10.5 10.5-.426 0-.847-.026-1.26-.075A8.5 8.5 0 1016.425 4.74c.05.413.075.833.075 1.259zm-1.732-2.04A9.08 9.08 0 0114.999 6a9 9 0 01-11.04 8.768l-.004-.002a9.367 9.367 0 01-.78-.218c-.393-.13-.8.21-.67.602a9.938 9.938 0 00.329.855l.004.01A10.002 10.002 0 0012 22a10.002 10.002 0 004.015-19.16l-.01-.005a9.745 9.745 0 00-.855-.328c-.392-.13-.732.276-.602.67a8.934 8.934 0 01.218.779l.002.005z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// MortarBoard returns a string representing an "mortar-board" Octicon.
func MortarBoard(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-mortar-board%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.693 1.066a.75.75 0 01.614 0l7.25 3.25a.75.75 0 010 1.368L13 6.831v2.794c0 1.024-.81 1.749-1.66 2.173-.893.447-2.075.702-3.34.702-.278 0-.55-.012-.816-.036a.75.75 0 01.133-1.494c.22.02.45.03.683.03 1.082 0 2.025-.221 2.67-.543.69-.345.83-.682.83-.832V7.503L8.307 8.934a.75.75 0 01-.614 0L4 7.28v1.663c.296.105.575.275.812.512.438.438.688 1.059.688 1.796v3a.75.75 0 01-.75.75h-3a.75.75 0 01-.75-.75v-3c0-.737.25-1.358.688-1.796.237-.237.516-.407.812-.512V6.606L.443 5.684a.75.75 0 010-1.368l7.25-3.25zM2.583 5L8 7.428 13.416 5 8 2.572 2.583 5zM2.5 11.25c0-.388.125-.611.25-.735a.704.704 0 01.5-.203c.19 0 .37.071.5.203.125.124.25.347.25.735v2.25H2.5v-2.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-mortar-board%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.292 2.06a.75.75 0 00-.584 0L.458 6.81a.75.75 0 000 1.38L4.25 9.793v3.803a2.901 2.901 0 00-1.327.757c-.579.58-.923 1.41-.923 2.43v4.5c0 .248.128.486.335.624.06.04.117.073.22.124.124.062.297.138.52.213.448.149 1.09.288 1.925.288s1.477-.14 1.925-.288c.223-.075.396-.15.52-.213a2.11 2.11 0 00.21-.117A.762.762 0 008 21.28v-4.5c0-1.018-.344-1.85-.923-2.428a2.9 2.9 0 00-1.327-.758v-3.17l5.958 2.516a.75.75 0 00.584 0l5.208-2.2v4.003a2.552 2.552 0 01-.079.085 4.057 4.057 0 01-.849.65c-.826.488-2.255 1.021-4.572 1.021-.612 0-1.162-.037-1.654-.1a.75.75 0 00-.192 1.487c.56.072 1.173.113 1.846.113 2.558 0 4.254-.592 5.334-1.23.538-.316.914-.64 1.163-.896a2.84 2.84 0 00.392-.482h.001A.75.75 0 0019 15v-4.892l4.542-1.917a.75.75 0 000-1.382l-11.25-4.75zM5 15c-.377 0-.745.141-1.017.413-.265.265-.483.7-.483 1.368v4.022c.299.105.797.228 1.5.228s1.201-.123 1.5-.228V16.78c0-.669-.218-1.103-.483-1.368A1.431 1.431 0 005 15zm7-3.564L2.678 7.5 12 3.564 21.322 7.5 12 11.436z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// MultiSelect returns a string representing an "multi-select" Octicon.
func MultiSelect(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-multi-select%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 2.5a.75.75 0 000 1.5h6.5a.75.75 0 000-1.5h-6.5zm4 5a.75.75 0 000 1.5h7.5a.75.75 0 000-1.5h-7.5zm0 5a.75.75 0 000 1.5h7.5a.75.75 0 000-1.5h-7.5zM3 8a1 1 0 11-2 0 1 1 0 012 0zm-1 6a1 1 0 100-2 1 1 0 000 2z\"></path><path d=\"M13.314 4.918L11.07 2.417A.25.25 0 0111.256 2h4.488a.25.25 0 01.186.417l-2.244 2.5a.25.25 0 01-.372 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-multi-select%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.75 5.5a.75.75 0 000 1.5h10a.75.75 0 000-1.5h-10zm5 6a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zm0 6a.75.75 0 000 1.5h11.5a.75.75 0 000-1.5H8.75zM5 12a1 1 0 11-2 0 1 1 0 012 0zm-1 7a1 1 0 100-2 1 1 0 000 2z\"></path><path d=\"M19.309 7.918l-2.245-2.501A.25.25 0 0117.25 5h4.49a.25.25 0 01.185.417l-2.244 2.5a.25.25 0 01-.372 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Mute returns a string representing an "mute" Octicon.
func Mute(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-mute%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 2.75a.75.75 0 00-1.238-.57L3.472 5H1.75A1.75 1.75 0 000 6.75v2.5C0 10.216.784 11 1.75 11h1.723l3.289 2.82A.75.75 0 008 13.25V2.75zM4.238 6.32L6.5 4.38v7.24L4.238 9.68a.75.75 0 00-.488-.18h-2a.25.25 0 01-.25-.25v-2.5a.25.25 0 01.25-.25h2a.75.75 0 00.488-.18zm7.042-1.1a.75.75 0 10-1.06 1.06L11.94 8l-1.72 1.72a.75.75 0 101.06 1.06L13 9.06l1.72 1.72a.75.75 0 101.06-1.06L14.06 8l1.72-1.72a.75.75 0 00-1.06-1.06L13 6.94l-1.72-1.72z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-mute%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 3.75a.75.75 0 00-1.255-.555L5.46 8H2.75A1.75 1.75 0 001 9.75v4.5c0 .966.784 1.75 1.75 1.75h2.71l5.285 4.805A.75.75 0 0012 20.25V3.75zM6.255 9.305l4.245-3.86v13.11l-4.245-3.86a.75.75 0 00-.505-.195h-3a.25.25 0 01-.25-.25v-4.5a.25.25 0 01.25-.25h3a.75.75 0 00.505-.195z\"></path><path d=\"M16.28 8.22a.75.75 0 10-1.06 1.06L17.94 12l-2.72 2.72a.75.75 0 101.06 1.06L19 13.06l2.72 2.72a.75.75 0 101.06-1.06L20.06 12l2.72-2.72a.75.75 0 00-1.06-1.06L19 10.94l-2.72-2.72z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// NoEntry returns a string representing an "no-entry" Octicon.
func NoEntry(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-no-entry%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M4.25 7.25a.75.75 0 000 1.5h7.5a.75.75 0 000-1.5h-7.5z\"></path><path fill-rule=\"evenodd\" d=\"M16 8A8 8 0 110 8a8 8 0 0116 0zm-1.5 0a6.5 6.5 0 11-13 0 6.5 6.5 0 0113 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-no-entry%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0zM12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zm6.25 11.75a.75.75 0 000-1.5H5.75a.75.75 0 000 1.5h12.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// NoEntryFill returns a string representing an "no-entry-fill" Octicon.
func NoEntryFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 12:
		return fmt.Sprintf("<svg class=\"octicon octicon-no-entry-fill%s\" height=\"12\" width=\"12\" viewbox=\"0 0 12 12\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 0a6 6 0 100 12A6 6 0 006 0zm3 5H3v2h6V5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// NorthStar returns a string representing an "north-star" Octicon.
func NorthStar(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-north-star%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8.5.75a.75.75 0 00-1.5 0v5.19L4.391 3.33a.75.75 0 10-1.06 1.061L5.939 7H.75a.75.75 0 000 1.5h5.19l-2.61 2.609a.75.75 0 101.061 1.06L7 9.561v5.189a.75.75 0 001.5 0V9.56l2.609 2.61a.75.75 0 101.06-1.061L9.561 8.5h5.189a.75.75 0 000-1.5H9.56l2.61-2.609a.75.75 0 00-1.061-1.06L8.5 5.939V.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-north-star%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12.5 1.25a.75.75 0 00-1.5 0v8.69L6.447 5.385a.75.75 0 10-1.061 1.06L9.94 11H1.25a.75.75 0 000 1.5h8.69l-4.554 4.553a.75.75 0 001.06 1.061L11 13.561v8.689a.75.75 0 001.5 0v-8.69l4.553 4.554a.75.75 0 001.061-1.06L13.561 12.5h8.689a.75.75 0 000-1.5h-8.69l4.554-4.553a.75.75 0 10-1.06-1.061L12.5 9.939V1.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Note returns a string representing an "note" Octicon.
func Note(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-note%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 3.75C0 2.784.784 2 1.75 2h12.5c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0114.25 14H1.75A1.75 1.75 0 010 12.25v-8.5zm1.75-.25a.25.25 0 00-.25.25v8.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25v-8.5a.25.25 0 00-.25-.25H1.75zM3.5 6.25a.75.75 0 01.75-.75h7a.75.75 0 010 1.5h-7a.75.75 0 01-.75-.75zm.75 2.25a.75.75 0 000 1.5h4a.75.75 0 000-1.5h-4z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-note%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 4.75C0 3.784.784 3 1.75 3h20.5c.966 0 1.75.784 1.75 1.75v14.5A1.75 1.75 0 0122.25 21H1.75A1.75 1.75 0 010 19.25V4.75zm1.75-.25a.25.25 0 00-.25.25v14.5c0 .138.112.25.25.25h20.5a.25.25 0 00.25-.25V4.75a.25.25 0 00-.25-.25H1.75z\"></path><path fill-rule=\"evenodd\" d=\"M5 8.75A.75.75 0 015.75 8h11.5a.75.75 0 010 1.5H5.75A.75.75 0 015 8.75zm0 4a.75.75 0 01.75-.75h5.5a.75.75 0 010 1.5h-5.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Number returns a string representing an "number" Octicon.
func Number(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-number%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.604.089A.75.75 0 016 .75v4.77h.711a.75.75 0 110 1.5H3.759a.75.75 0 110-1.5H4.5V2.15l-.334.223a.75.75 0 01-.832-1.248l1.5-1a.75.75 0 01.77-.037zM9 4.75A.75.75 0 019.75 4h4a.75.75 0 01.53 1.28l-1.89 1.892c.312.076.604.18.867.319.742.391 1.244 1.063 1.244 2.005 0 .653-.231 1.208-.629 1.627-.386.408-.894.653-1.408.777-1.01.243-2.225.063-3.124-.527a.75.75 0 01.822-1.254c.534.35 1.32.474 1.951.322.306-.073.53-.201.67-.349.129-.136.218-.32.218-.596 0-.308-.123-.509-.444-.678-.373-.197-.98-.318-1.806-.318a.75.75 0 01-.53-1.28l1.72-1.72H9.75A.75.75 0 019 4.75zm-3.587 5.763c-.35-.05-.77.113-.983.572a.75.75 0 11-1.36-.632c.508-1.094 1.589-1.565 2.558-1.425 1 .145 1.872.945 1.872 2.222 0 1.433-1.088 2.192-1.79 2.681-.308.216-.571.397-.772.573H7a.75.75 0 010 1.5H3.75a.75.75 0 01-.75-.75c0-.69.3-1.211.67-1.61.348-.372.8-.676 1.15-.92.8-.56 1.18-.904 1.18-1.474 0-.473-.267-.69-.587-.737z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-number%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.114 2.094a.75.75 0 01.386.656V9h1.252a.75.75 0 110 1.5H5.75a.75.75 0 010-1.5H7V4.103l-.853.533a.75.75 0 01-.795-1.272l2-1.25a.75.75 0 01.762-.02zm4.889 5.66a.75.75 0 01.75-.75h5.232a.75.75 0 01.53 1.28l-2.776 2.777c.55.097 1.057.253 1.492.483.905.477 1.504 1.284 1.504 2.418 0 .966-.471 1.75-1.172 2.27-.687.511-1.587.77-2.521.77-1.367 0-2.274-.528-2.667-.756a.75.75 0 01.755-1.297c.331.193.953.553 1.912.553.673 0 1.243-.188 1.627-.473.37-.275.566-.635.566-1.067 0-.5-.219-.836-.703-1.091-.538-.284-1.375-.443-2.471-.443a.75.75 0 01-.53-1.28l2.643-2.644h-3.421a.75.75 0 01-.75-.75zM7.88 15.215a1.4 1.4 0 00-1.446.83.75.75 0 01-1.37-.61 2.9 2.9 0 012.986-1.71 2.565 2.565 0 011.557.743c.434.446.685 1.058.685 1.778 0 1.641-1.254 2.437-2.12 2.986-.538.341-1.18.694-1.495 1.273H9.75a.75.75 0 010 1.5h-4a.75.75 0 01-.75-.75c0-1.799 1.337-2.63 2.243-3.21 1.032-.659 1.55-1.031 1.55-1.8 0-.355-.116-.584-.26-.732a1.068 1.068 0 00-.652-.298z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Organization returns a string representing an "organization" Octicon.
func Organization(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-organization%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 14.25c0 .138.112.25.25.25H4v-1.25a.75.75 0 01.75-.75h2.5a.75.75 0 01.75.75v1.25h2.25a.25.25 0 00.25-.25V1.75a.25.25 0 00-.25-.25h-8.5a.25.25 0 00-.25.25v12.5zM1.75 16A1.75 1.75 0 010 14.25V1.75C0 .784.784 0 1.75 0h8.5C11.216 0 12 .784 12 1.75v12.5c0 .085-.006.168-.018.25h2.268a.25.25 0 00.25-.25V8.285a.25.25 0 00-.111-.208l-1.055-.703a.75.75 0 11.832-1.248l1.055.703c.487.325.779.871.779 1.456v5.965A1.75 1.75 0 0114.25 16h-3.5a.75.75 0 01-.197-.026c-.099.017-.2.026-.303.026h-3a.75.75 0 01-.75-.75V14h-1v1.25a.75.75 0 01-.75.75h-3zM3 3.75A.75.75 0 013.75 3h.5a.75.75 0 010 1.5h-.5A.75.75 0 013 3.75zM3.75 6a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM3 9.75A.75.75 0 013.75 9h.5a.75.75 0 010 1.5h-.5A.75.75 0 013 9.75zM7.75 9a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM7 6.75A.75.75 0 017.75 6h.5a.75.75 0 010 1.5h-.5A.75.75 0 017 6.75zM7.75 3a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-organization%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M6.25 12a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM5.5 9.25a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5a.75.75 0 01-.75-.75zM6.25 5a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM9 12.75a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5a.75.75 0 01-.75-.75zm.75-4.25a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM9 5.75A.75.75 0 019.75 5h.5a.75.75 0 010 1.5h-.5A.75.75 0 019 5.75zM13.25 12a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zm-.75-2.75a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5a.75.75 0 01-.75-.75zM13.25 5a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5z\"></path><path fill-rule=\"evenodd\" d=\"M2 20a2 2 0 002 2h3.75a.75.75 0 00.75-.75V19h3v2.25c0 .414.336.75.75.75H16c.092 0 .183-.006.272-.018a.753.753 0 00.166.018H20a2 2 0 002-2v-8a2 2 0 00-.8-1.6l-.5-.375a.75.75 0 00-.9 1.2l.5.375a.5.5 0 01.2.4v8a.5.5 0 01-.5.5h-2.063c.041-.16.063-.327.063-.5V3a2 2 0 00-2-2H4a2 2 0 00-2 2v17zm2 .5a.5.5 0 01-.5-.5V3a.5.5 0 01.5-.5h12a.5.5 0 01.5.5v17a.5.5 0 01-.5.5h-3v-2.25a.75.75 0 00-.75-.75h-4.5a.75.75 0 00-.75.75v2.25H4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Package returns a string representing an "package" Octicon.
func Package(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-package%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.878.392a1.75 1.75 0 00-1.756 0l-5.25 3.045A1.75 1.75 0 001 4.951v6.098c0 .624.332 1.2.872 1.514l5.25 3.045a1.75 1.75 0 001.756 0l5.25-3.045c.54-.313.872-.89.872-1.514V4.951c0-.624-.332-1.2-.872-1.514L8.878.392zM7.875 1.69a.25.25 0 01.25 0l4.63 2.685L8 7.133 3.245 4.375l4.63-2.685zM2.5 5.677v5.372c0 .09.047.171.125.216l4.625 2.683V8.432L2.5 5.677zm6.25 8.271l4.625-2.683a.25.25 0 00.125-.216V5.677L8.75 8.432v5.516z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-package%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.876.64a1.75 1.75 0 00-1.75 0l-8.25 4.762a1.75 1.75 0 00-.875 1.515v9.525c0 .625.334 1.203.875 1.515l8.25 4.763a1.75 1.75 0 001.75 0l8.25-4.762a1.75 1.75 0 00.875-1.516V6.917a1.75 1.75 0 00-.875-1.515L12.876.639zm-1 1.298a.25.25 0 01.25 0l7.625 4.402-7.75 4.474-7.75-4.474 7.625-4.402zM3.501 7.64v8.803c0 .09.048.172.125.216l7.625 4.402v-8.947L3.501 7.64zm9.25 13.421l7.625-4.402a.25.25 0 00.125-.216V7.639l-7.75 4.474v8.947z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// PackageDependencies returns a string representing an "package-dependencies" Octicon.
func PackageDependencies(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-package-dependencies%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.122.392a1.75 1.75 0 011.756 0l5.25 3.045c.54.313.872.89.872 1.514V7.25a.75.75 0 01-1.5 0V5.677L7.75 8.432v6.384a1 1 0 01-1.502.865L.872 12.563A1.75 1.75 0 010 11.049V4.951c0-.624.332-1.2.872-1.514L6.122.392zM7.125 1.69l4.63 2.685L7 7.133 2.245 4.375l4.63-2.685a.25.25 0 01.25 0zM1.5 11.049V5.677l4.75 2.755v5.516l-4.625-2.683a.25.25 0 01-.125-.216zm11.672-.282a.75.75 0 10-1.087-1.034l-2.378 2.5a.75.75 0 000 1.034l2.378 2.5a.75.75 0 101.087-1.034L11.999 13.5h3.251a.75.75 0 000-1.5h-3.251l1.173-1.233z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-package-dependencies%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.126.64a1.75 1.75 0 011.75 0l8.25 4.762c.103.06.199.128.286.206a.748.748 0 01.554.96c.023.113.035.23.035.35v3.332a.75.75 0 01-1.5 0V7.64l-7.75 4.474V22.36a.75.75 0 01-1.125.65l-8.75-5.052a1.75 1.75 0 01-.875-1.515V6.917c0-.119.012-.236.035-.35a.748.748 0 01.554-.96 1.75 1.75 0 01.286-.205L9.126.639zM1.501 7.638v8.803c0 .09.048.172.125.216l7.625 4.402v-8.947l-7.75-4.474zm8.5 3.175L2.251 6.34l7.625-4.402a.25.25 0 01.25 0l7.625 4.402-7.75 4.474z\"></path><path d=\"M16.617 17.5l2.895-2.702a.75.75 0 00-1.024-1.096l-4.285 4a.75.75 0 000 1.096l4.285 4a.75.75 0 101.024-1.096L16.617 19h6.633a.75.75 0 000-1.5h-6.633z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// PackageDependents returns a string representing an "package-dependents" Octicon.
func PackageDependents(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-package-dependents%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.122.392a1.75 1.75 0 011.756 0l5.25 3.045c.54.313.872.89.872 1.514V7.25a.75.75 0 01-1.5 0V5.677L7.75 8.432v6.384a1 1 0 01-1.502.865L.872 12.563A1.75 1.75 0 010 11.049V4.951c0-.624.332-1.2.872-1.514L6.122.392zM7.125 1.69l4.63 2.685L7 7.133 2.245 4.375l4.63-2.685a.25.25 0 01.25 0zM1.5 11.049V5.677l4.75 2.755v5.516l-4.625-2.683a.25.25 0 01-.125-.216zm10.828 3.684a.75.75 0 101.087 1.034l2.378-2.5a.75.75 0 000-1.034l-2.378-2.5a.75.75 0 00-1.087 1.034L13.501 12H10.25a.75.75 0 000 1.5h3.251l-1.173 1.233z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-package-dependents%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.126.64a1.75 1.75 0 011.75 0l8.25 4.762c.103.06.199.128.286.206a.748.748 0 01.554.96c.023.113.035.23.035.35v3.332a.75.75 0 01-1.5 0V7.64l-7.75 4.474V22.36a.75.75 0 01-1.125.65l-8.75-5.052a1.75 1.75 0 01-.875-1.515V6.917c0-.119.012-.236.035-.35a.748.748 0 01.554-.96 1.75 1.75 0 01.286-.205L9.126.639zM1.501 7.638v8.803c0 .09.048.172.125.216l7.625 4.402v-8.947l-7.75-4.474zm8.5 3.175L2.251 6.34l7.625-4.402a.25.25 0 01.25 0l7.625 4.402-7.75 4.474z\"></path><path d=\"M21.347 17.5l-2.894-2.702a.75.75 0 111.023-1.096l4.286 4a.75.75 0 010 1.096l-4.286 4a.75.75 0 11-1.023-1.096L21.347 19h-6.633a.75.75 0 010-1.5h6.633z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Paintbrush returns a string representing an "paintbrush" Octicon.
func Paintbrush(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-paintbrush%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.134 1.535C9.722 2.562 8.16 4.057 6.889 5.312 5.8 6.387 5.041 7.401 4.575 8.294a3.745 3.745 0 00-3.227 1.054c-.43.431-.69 1.066-.86 1.657a11.982 11.982 0 00-.358 1.914A21.263 21.263 0 000 15.203v.054l.75-.007-.007.75h.054a14.404 14.404 0 00.654-.012 21.243 21.243 0 001.63-.118c.62-.07 1.3-.18 1.914-.357.592-.17 1.226-.43 1.657-.861a3.745 3.745 0 001.055-3.217c.908-.461 1.942-1.216 3.04-2.3 1.279-1.262 2.764-2.825 3.775-4.249.501-.706.923-1.428 1.125-2.096.2-.659.235-1.469-.368-2.07-.606-.607-1.42-.55-2.069-.34-.66.213-1.376.646-2.076 1.155zm-3.95 8.48a3.76 3.76 0 00-1.19-1.192 9.758 9.758 0 011.161-1.607l1.658 1.658a9.853 9.853 0 01-1.63 1.142zM.742 16l.007-.75-.75.008A.75.75 0 00.743 16zM12.016 2.749c-1.224.89-2.605 2.189-3.822 3.384l1.718 1.718c1.21-1.205 2.51-2.597 3.387-3.833.47-.662.78-1.227.912-1.662.134-.444.032-.551.009-.575h-.001V1.78c-.014-.014-.112-.113-.548.027-.432.14-.995.462-1.655.942zM1.62 13.089a19.56 19.56 0 00-.104 1.395 19.55 19.55 0 001.396-.104 10.528 10.528 0 001.668-.309c.526-.151.856-.325 1.011-.48a2.25 2.25 0 00-3.182-3.182c-.155.155-.329.485-.48 1.01a10.515 10.515 0 00-.309 1.67z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// PaperAirplane returns a string representing an "paper-airplane" Octicon.
func PaperAirplane(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-paper-airplane%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.592 2.712L2.38 7.25h4.87a.75.75 0 110 1.5H2.38l-.788 4.538L13.929 8 1.592 2.712zM.989 8L.064 2.68a1.341 1.341 0 011.85-1.462l13.402 5.744a1.13 1.13 0 010 2.076L1.913 14.782a1.341 1.341 0 01-1.85-1.463L.99 8z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-paper-airplane%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.513 1.96a1.374 1.374 0 011.499-.21l19.335 9.215a1.146 1.146 0 010 2.07L3.012 22.25a1.374 1.374 0 01-1.947-1.46L2.49 12 1.065 3.21a1.374 1.374 0 01.448-1.25zm2.375 10.79l-1.304 8.042L21.031 12 2.584 3.208l1.304 8.042h7.362a.75.75 0 010 1.5H3.888z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Paste returns a string representing an "paste" Octicon.
func Paste(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-paste%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.75 1a.75.75 0 00-.75.75v3c0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75v-3a.75.75 0 00-.75-.75h-4.5zm.75 3V2.5h3V4h-3zm-2.874-.467a.75.75 0 00-.752-1.298A1.75 1.75 0 002 3.75v9.5c0 .966.784 1.75 1.75 1.75h8.5A1.75 1.75 0 0014 13.25v-9.5a1.75 1.75 0 00-.874-1.515.75.75 0 10-.752 1.298.25.25 0 01.126.217v9.5a.25.25 0 01-.25.25h-8.5a.25.25 0 01-.25-.25v-9.5a.25.25 0 01.126-.217z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-paste%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.962 2.513a.75.75 0 01-.475.949l-.816.272a.25.25 0 00-.171.237V21.25c0 .138.112.25.25.25h14.5a.25.25 0 00.25-.25V3.97a.25.25 0 00-.17-.236l-.817-.272a.75.75 0 01.474-1.424l.816.273A1.75 1.75 0 0121 3.97v17.28A1.75 1.75 0 0119.25 23H4.75A1.75 1.75 0 013 21.25V3.97a1.75 1.75 0 011.197-1.66l.816-.272a.75.75 0 01.949.475z\"></path><path fill-rule=\"evenodd\" d=\"M7 1.75C7 .784 7.784 0 8.75 0h6.5C16.216 0 17 .784 17 1.75v1.5A1.75 1.75 0 0115.25 5h-6.5A1.75 1.75 0 017 3.25v-1.5zm1.75-.25a.25.25 0 00-.25.25v1.5c0 .138.112.25.25.25h6.5a.25.25 0 00.25-.25v-1.5a.25.25 0 00-.25-.25h-6.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Pencil returns a string representing an "pencil" Octicon.
func Pencil(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-pencil%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.013 1.427a1.75 1.75 0 012.474 0l1.086 1.086a1.75 1.75 0 010 2.474l-8.61 8.61c-.21.21-.47.364-.756.445l-3.251.93a.75.75 0 01-.927-.928l.929-3.25a1.75 1.75 0 01.445-.758l8.61-8.61zm1.414 1.06a.25.25 0 00-.354 0L10.811 3.75l1.439 1.44 1.263-1.263a.25.25 0 000-.354l-1.086-1.086zM11.189 6.25L9.75 4.81l-6.286 6.287a.25.25 0 00-.064.108l-.558 1.953 1.953-.558a.249.249 0 00.108-.064l6.286-6.286z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-pencil%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M17.263 2.177a1.75 1.75 0 012.474 0l2.586 2.586a1.75 1.75 0 010 2.474L19.53 10.03l-.012.013L8.69 20.378a1.75 1.75 0 01-.699.409l-5.523 1.68a.75.75 0 01-.935-.935l1.673-5.5a1.75 1.75 0 01.466-.756L14.476 4.963l2.787-2.786zm-2.275 4.371l-10.28 9.813a.25.25 0 00-.067.108l-1.264 4.154 4.177-1.271a.25.25 0 00.1-.059l10.273-9.806-2.94-2.939zM19 8.44l2.263-2.262a.25.25 0 000-.354l-2.586-2.586a.25.25 0 00-.354 0L16.061 5.5 19 8.44z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// People returns a string representing an "people" Octicon.
func People(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-people%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.5 3.5a2 2 0 100 4 2 2 0 000-4zM2 5.5a3.5 3.5 0 115.898 2.549 5.507 5.507 0 013.034 4.084.75.75 0 11-1.482.235 4.001 4.001 0 00-7.9 0 .75.75 0 01-1.482-.236A5.507 5.507 0 013.102 8.05 3.49 3.49 0 012 5.5zM11 4a.75.75 0 100 1.5 1.5 1.5 0 01.666 2.844.75.75 0 00-.416.672v.352a.75.75 0 00.574.73c1.2.289 2.162 1.2 2.522 2.372a.75.75 0 101.434-.44 5.01 5.01 0 00-2.56-3.012A3 3 0 0011 4z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-people%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.5 8a5.5 5.5 0 118.596 4.547 9.005 9.005 0 015.9 8.18.75.75 0 01-1.5.045 7.5 7.5 0 00-14.993 0 .75.75 0 01-1.499-.044 9.005 9.005 0 015.9-8.181A5.494 5.494 0 013.5 8zM9 4a4 4 0 100 8 4 4 0 000-8z\"></path><path d=\"M17.29 8c-.148 0-.292.01-.434.03a.75.75 0 11-.212-1.484 4.53 4.53 0 013.38 8.097 6.69 6.69 0 013.956 6.107.75.75 0 01-1.5 0 5.193 5.193 0 00-3.696-4.972l-.534-.16v-1.676l.41-.209A3.03 3.03 0 0017.29 8z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Person returns a string representing an "person" Octicon.
func Person(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-person%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.5 5a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0zm.061 3.073a4 4 0 10-5.123 0 6.004 6.004 0 00-3.431 5.142.75.75 0 001.498.07 4.5 4.5 0 018.99 0 .75.75 0 101.498-.07 6.005 6.005 0 00-3.432-5.142z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-person%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 2.5a5.5 5.5 0 00-3.096 10.047 9.005 9.005 0 00-5.9 8.18.75.75 0 001.5.045 7.5 7.5 0 0114.993 0 .75.75 0 101.499-.044 9.005 9.005 0 00-5.9-8.181A5.5 5.5 0 0012 2.5zM8 8a4 4 0 118 0 4 4 0 01-8 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// PersonAdd returns a string representing an "person-add" Octicon.
func PersonAdd(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-person-add%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13.25 0a.75.75 0 01.75.75V2h1.25a.75.75 0 010 1.5H14v1.25a.75.75 0 01-1.5 0V3.5h-1.25a.75.75 0 010-1.5h1.25V.75a.75.75 0 01.75-.75zM5.5 4a2 2 0 100 4 2 2 0 000-4zm2.4 4.548a3.5 3.5 0 10-4.799 0 5.527 5.527 0 00-3.1 4.66.75.75 0 101.498.085A4.01 4.01 0 015.5 9.5a4.01 4.01 0 014.001 3.793.75.75 0 101.498-.086 5.527 5.527 0 00-3.1-4.659z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-person-add%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M19.25 1a.75.75 0 01.75.75V4h2.25a.75.75 0 010 1.5H20v2.25a.75.75 0 01-1.5 0V5.5h-2.25a.75.75 0 010-1.5h2.25V1.75a.75.75 0 01.75-.75zM9 6a3.5 3.5 0 100 7 3.5 3.5 0 000-7zM4 9.5a5 5 0 117.916 4.062 7.973 7.973 0 015.018 7.166.75.75 0 11-1.499.044 6.469 6.469 0 00-12.932 0 .75.75 0 01-1.499-.044 7.973 7.973 0 015.059-7.181A4.993 4.993 0 014 9.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// PersonFill returns a string representing an "person-fill" Octicon.
func PersonFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-person-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M4.243 4.757a3.757 3.757 0 115.851 3.119 6.006 6.006 0 013.9 5.339.75.75 0 01-.715.784H2.721a.75.75 0 01-.714-.784 6.006 6.006 0 013.9-5.34 3.753 3.753 0 01-1.664-3.118z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-person-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12 2.5a5.25 5.25 0 00-2.519 9.857 9.005 9.005 0 00-6.477 8.37.75.75 0 00.727.773H20.27a.75.75 0 00.727-.772 9.005 9.005 0 00-6.477-8.37A5.25 5.25 0 0012 2.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Pin returns a string representing an "pin" Octicon.
func Pin(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-pin%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.456.734a1.75 1.75 0 012.826.504l.613 1.327a3.081 3.081 0 002.084 1.707l2.454.584c1.332.317 1.8 1.972.832 2.94L11.06 10l3.72 3.72a.75.75 0 11-1.061 1.06L10 11.06l-2.204 2.205c-.968.968-2.623.5-2.94-.832l-.584-2.454a3.081 3.081 0 00-1.707-2.084l-1.327-.613a1.75 1.75 0 01-.504-2.826L4.456.734zM5.92 1.866a.25.25 0 00-.404-.072L1.794 5.516a.25.25 0 00.072.404l1.328.613A4.582 4.582 0 015.73 9.63l.584 2.454a.25.25 0 00.42.12l5.47-5.47a.25.25 0 00-.12-.42L9.63 5.73a4.581 4.581 0 01-3.098-2.537L5.92 1.866z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-pin%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.886 1.553a1.75 1.75 0 012.869.604l.633 1.629a5.666 5.666 0 003.725 3.395l3.959 1.131a1.75 1.75 0 01.757 2.92L16.06 15l5.594 5.595a.75.75 0 11-1.06 1.06L15 16.061l-3.768 3.768a1.75 1.75 0 01-2.92-.757l-1.131-3.96a5.667 5.667 0 00-3.395-3.724l-1.63-.633a1.75 1.75 0 01-.603-2.869l6.333-6.333zm6.589 12.912l-.005.005-.005.005-4.294 4.293a.25.25 0 01-.417-.108l-1.13-3.96A7.166 7.166 0 004.33 9.99L2.7 9.356a.25.25 0 01-.086-.41l6.333-6.332a.25.25 0 01.41.086l.633 1.63a7.167 7.167 0 004.71 4.293l3.96 1.131a.25.25 0 01.108.417l-4.293 4.294z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Play returns a string representing an "play" Octicon.
func Play(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-play%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 8a6.5 6.5 0 1113 0 6.5 6.5 0 01-13 0zM8 0a8 8 0 100 16A8 8 0 008 0zM6.379 5.227A.25.25 0 006 5.442v5.117a.25.25 0 00.379.214l4.264-2.559a.25.25 0 000-.428L6.379 5.227z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-play%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M9.5 15.584V8.416a.5.5 0 01.77-.42l5.576 3.583a.5.5 0 010 .842l-5.576 3.584a.5.5 0 01-.77-.42z\"></path><path fill-rule=\"evenodd\" d=\"M12 2.5a9.5 9.5 0 100 19 9.5 9.5 0 000-19zM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11-4.925 11-11 11S1 18.075 1 12z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Plug returns a string representing an "plug" Octicon.
func Plug(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-plug%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.276 3.09a.25.25 0 01.192-.09h.782a.25.25 0 01.25.25v8.5a.25.25 0 01-.25.25h-.782a.25.25 0 01-.192-.09l-.95-1.14a.75.75 0 00-.483-.264l-3.124-.39a.25.25 0 01-.219-.249V5.133a.25.25 0 01.219-.248l3.124-.39a.75.75 0 00.483-.265l.95-1.14zM4 8v1.867a1.75 1.75 0 001.533 1.737l2.83.354.761.912c.332.4.825.63 1.344.63h.782A1.75 1.75 0 0013 11.75V11h2.25a.75.75 0 000-1.5H13v-4h2.25a.75.75 0 000-1.5H13v-.75a1.75 1.75 0 00-1.75-1.75h-.782c-.519 0-1.012.23-1.344.63l-.76.913-2.831.353A1.75 1.75 0 004 5.133V6.5H2.5A2.5 2.5 0 000 9v5.25a.75.75 0 001.5 0V9a1 1 0 011-1H4z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-plug%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7 11.5v3.848a1.75 1.75 0 001.57 1.74l6.055.627 1.006 1.174a1.75 1.75 0 001.329.611h1.29A1.75 1.75 0 0020 17.75V15.5h3.25a.75.75 0 000-1.5H20V7.5h3.25a.75.75 0 000-1.5H20V3.75A1.75 1.75 0 0018.25 2h-1.29c-.51 0-.996.223-1.329.611l-1.006 1.174-6.055.626A1.75 1.75 0 007 6.151V10H2.937A2.938 2.938 0 000 12.938v8.312a.75.75 0 001.5 0v-8.313c0-.793.644-1.437 1.438-1.437H7zm9.77-7.913a.25.25 0 01.19-.087h1.29a.25.25 0 01.25.25v14a.25.25 0 01-.25.25h-1.29a.25.25 0 01-.19-.087l-1.2-1.401a.75.75 0 00-.493-.258l-6.353-.657a.25.25 0 01-.224-.249V6.152a.25.25 0 01.224-.249l6.353-.657a.75.75 0 00.492-.258l1.201-1.4z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Plus returns a string representing an "plus" Octicon.
func Plus(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-plus%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.75 2a.75.75 0 01.75.75V7h4.25a.75.75 0 110 1.5H8.5v4.25a.75.75 0 11-1.5 0V8.5H2.75a.75.75 0 010-1.5H7V2.75A.75.75 0 017.75 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-plus%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.75 4.5a.75.75 0 01.75.75V11h5.75a.75.75 0 010 1.5H12.5v5.75a.75.75 0 01-1.5 0V12.5H5.25a.75.75 0 010-1.5H11V5.25a.75.75 0 01.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// PlusCircle returns a string representing an "plus-circle" Octicon.
func PlusCircle(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-plus-circle%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 8a6.5 6.5 0 1113 0 6.5 6.5 0 01-13 0zM8 0a8 8 0 100 16A8 8 0 008 0zm.75 4.75a.75.75 0 00-1.5 0v2.5h-2.5a.75.75 0 000 1.5h2.5v2.5a.75.75 0 001.5 0v-2.5h2.5a.75.75 0 000-1.5h-2.5v-2.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-plus-circle%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12.75 7.75a.75.75 0 00-1.5 0v3.5h-3.5a.75.75 0 000 1.5h3.5v3.5a.75.75 0 001.5 0v-3.5h3.5a.75.75 0 000-1.5h-3.5v-3.5z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Project returns a string representing an "project" Octicon.
func Project(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-project%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 0A1.75 1.75 0 000 1.75v12.5C0 15.216.784 16 1.75 16h12.5A1.75 1.75 0 0016 14.25V1.75A1.75 1.75 0 0014.25 0H1.75zM1.5 1.75a.25.25 0 01.25-.25h12.5a.25.25 0 01.25.25v12.5a.25.25 0 01-.25.25H1.75a.25.25 0 01-.25-.25V1.75zM11.75 3a.75.75 0 00-.75.75v7.5a.75.75 0 001.5 0v-7.5a.75.75 0 00-.75-.75zm-8.25.75a.75.75 0 011.5 0v5.5a.75.75 0 01-1.5 0v-5.5zM8 3a.75.75 0 00-.75.75v3.5a.75.75 0 001.5 0v-3.5A.75.75 0 008 3z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-project%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M7.25 6a.75.75 0 00-.75.75v7.5a.75.75 0 001.5 0v-7.5A.75.75 0 007.25 6zM12 6a.75.75 0 00-.75.75v4.5a.75.75 0 001.5 0v-4.5A.75.75 0 0012 6zm4 .75a.75.75 0 011.5 0v9.5a.75.75 0 01-1.5 0v-9.5z\"></path><path fill-rule=\"evenodd\" d=\"M3.75 2A1.75 1.75 0 002 3.75v16.5c0 .966.784 1.75 1.75 1.75h16.5A1.75 1.75 0 0022 20.25V3.75A1.75 1.75 0 0020.25 2H3.75zM3.5 3.75a.25.25 0 01.25-.25h16.5a.25.25 0 01.25.25v16.5a.25.25 0 01-.25.25H3.75a.25.25 0 01-.25-.25V3.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Pulse returns a string representing an "pulse" Octicon.
func Pulse(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-pulse%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 2a.75.75 0 01.696.471L10 10.731l1.304-3.26A.75.75 0 0112 7h3.25a.75.75 0 010 1.5h-2.742l-1.812 4.528a.75.75 0 01-1.392 0L6 4.77 4.696 8.03A.75.75 0 014 8.5H.75a.75.75 0 010-1.5h2.742l1.812-4.529A.75.75 0 016 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-pulse%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.002 2.5a.75.75 0 01.691.464l6.302 15.305 2.56-6.301a.75.75 0 01.695-.468h4a.75.75 0 010 1.5h-3.495l-3.06 7.532a.75.75 0 01-1.389.004L8.997 5.21l-3.054 7.329A.75.75 0 015.25 13H.75a.75.75 0 010-1.5h4l3.558-8.538a.75.75 0 01.694-.462z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Question returns a string representing an "question" Octicon.
func Question(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-question%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 1.5a6.5 6.5 0 100 13 6.5 6.5 0 000-13zM0 8a8 8 0 1116 0A8 8 0 010 8zm9 3a1 1 0 11-2 0 1 1 0 012 0zM6.92 6.085c.081-.16.19-.299.34-.398.145-.097.371-.187.74-.187.28 0 .553.087.738.225A.613.613 0 019 6.25c0 .177-.04.264-.077.318a.956.956 0 01-.277.245c-.076.051-.158.1-.258.161l-.007.004a7.728 7.728 0 00-.313.195 2.416 2.416 0 00-.692.661.75.75 0 001.248.832.956.956 0 01.276-.245 6.3 6.3 0 01.26-.16l.006-.004c.093-.057.204-.123.313-.195.222-.149.487-.355.692-.662.214-.32.329-.702.329-1.15 0-.76-.36-1.348-.863-1.725A2.76 2.76 0 008 4c-.631 0-1.155.16-1.572.438-.413.276-.68.638-.849.977a.75.75 0 101.342.67z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-question%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M10.97 8.265a1.45 1.45 0 00-.487.57.75.75 0 01-1.341-.67c.2-.402.513-.826.997-1.148C10.627 6.69 11.244 6.5 12 6.5c.658 0 1.369.195 1.934.619a2.45 2.45 0 011.004 2.006c0 1.033-.513 1.72-1.027 2.215-.19.183-.399.358-.579.508l-.147.123a4.329 4.329 0 00-.435.409v1.37a.75.75 0 11-1.5 0v-1.473c0-.237.067-.504.247-.736.22-.28.486-.517.718-.714l.183-.153.001-.001c.172-.143.324-.27.47-.412.368-.355.569-.676.569-1.136a.953.953 0 00-.404-.806C12.766 8.118 12.384 8 12 8c-.494 0-.814.121-1.03.265zM13 17a1 1 0 11-2 0 1 1 0 012 0z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Quote returns a string representing an "quote" Octicon.
func Quote(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-quote%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 2.5a.75.75 0 000 1.5h10.5a.75.75 0 000-1.5H1.75zm4 5a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5zm0 5a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5zM2.5 7.75a.75.75 0 00-1.5 0v6a.75.75 0 001.5 0v-6z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-quote%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3 6.25a.75.75 0 01.75-.75h13.5a.75.75 0 010 1.5H3.75A.75.75 0 013 6.25zM3.75 11a.75.75 0 01.75.75v7a.75.75 0 01-1.5 0v-7a.75.75 0 01.75-.75zM8 12.313a.75.75 0 01.75-.75h11.5a.75.75 0 010 1.5H8.75a.75.75 0 01-.75-.75zm0 5.937a.75.75 0 01.75-.75h11.5a.75.75 0 010 1.5H8.75a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Reply returns a string representing an "reply" Octicon.
func Reply(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-reply%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.78 1.97a.75.75 0 010 1.06L3.81 6h6.44A4.75 4.75 0 0115 10.75v2.5a.75.75 0 01-1.5 0v-2.5a3.25 3.25 0 00-3.25-3.25H3.81l2.97 2.97a.75.75 0 11-1.06 1.06L1.47 7.28a.75.75 0 010-1.06l4.25-4.25a.75.75 0 011.06 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-reply%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M10.53 5.03a.75.75 0 10-1.06-1.06l-6.25 6.25a.75.75 0 000 1.06l6.25 6.25a.75.75 0 101.06-1.06L5.56 11.5H17a3.248 3.248 0 013.25 3.248v4.502a.75.75 0 001.5 0v-4.502A4.748 4.748 0 0017 10H5.56l4.97-4.97z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Repo returns a string representing an "repo" Octicon.
func Repo(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 2.5A2.5 2.5 0 014.5 0h8.75a.75.75 0 01.75.75v12.5a.75.75 0 01-.75.75h-2.5a.75.75 0 110-1.5h1.75v-2h-8a1 1 0 00-.714 1.7.75.75 0 01-1.072 1.05A2.495 2.495 0 012 11.5v-9zm10.5-1V9h-8c-.356 0-.694.074-1 .208V2.5a1 1 0 011-1h8zM5 12.25v3.25a.25.25 0 00.4.2l1.45-1.087a.25.25 0 01.3 0L8.6 15.7a.25.25 0 00.4-.2v-3.25a.25.25 0 00-.25-.25h-3.5a.25.25 0 00-.25.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3 2.75A2.75 2.75 0 015.75 0h14.5a.75.75 0 01.75.75v20.5a.75.75 0 01-.75.75h-6a.75.75 0 010-1.5h5.25v-4H6A1.5 1.5 0 004.5 18v.75c0 .716.43 1.334 1.05 1.605a.75.75 0 01-.6 1.374A3.25 3.25 0 013 18.75v-16zM19.5 1.5V15H6c-.546 0-1.059.146-1.5.401V2.75c0-.69.56-1.25 1.25-1.25H19.5z\"></path><path d=\"M7 18.25a.25.25 0 01.25-.25h5a.25.25 0 01.25.25v5.01a.25.25 0 01-.397.201l-2.206-1.604a.25.25 0 00-.294 0L7.397 23.46a.25.25 0 01-.397-.2v-5.01z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// RepoClone returns a string representing an "repo-clone" Octicon.
func RepoClone(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-clone%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M15 0H9v7c0 .55.45 1 1 1h1v1h1V8h3c.55 0 1-.45 1-1V1c0-.55-.45-1-1-1zm-4 7h-1V6h1v1zm4 0h-3V6h3v1zm0-2h-4V1h4v4zM4 5H3V4h1v1zm0-2H3V2h1v1zM2 1h6V0H1C.45 0 0 .45 0 1v12c0 .55.45 1 1 1h2v2l1.5-1.5L6 16v-2h5c.55 0 1-.45 1-1v-3H2V1zm9 10v2H6v-1H3v1H1v-2h10zM3 8h1v1H3V8zm1-1H3V6h1v1z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// RepoDeleted returns a string representing an "repo-deleted" Octicon.
func RepoDeleted(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-deleted%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M1 2.5A2.5 2.5 0 013.5 0h8.75a.75.75 0 01.75.75v7.5a.75.75 0 01-1.5 0V1.5h-8a1 1 0 00-1 1v6.708A2.492 2.492 0 013.5 9h4.75a.75.75 0 010 1.5H3.5a1 1 0 100 2h4.75a.75.75 0 010 1.5H3.5A2.5 2.5 0 011 11.5v-9z\"></path><path d=\"M11.28 10.22a.75.75 0 10-1.06 1.06L11.94 13l-1.72 1.72a.75.75 0 101.06 1.06L13 14.06l1.72 1.72a.75.75 0 101.06-1.06L14.06 13l1.72-1.72a.75.75 0 10-1.06-1.06L13 11.94l-1.72-1.72z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// RepoForked returns a string representing an "repo-forked" Octicon.
func RepoForked(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-forked%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5 3.25a.75.75 0 11-1.5 0 .75.75 0 011.5 0zm0 2.122a2.25 2.25 0 10-1.5 0v.878A2.25 2.25 0 005.75 8.5h1.5v2.128a2.251 2.251 0 101.5 0V8.5h1.5a2.25 2.25 0 002.25-2.25v-.878a2.25 2.25 0 10-1.5 0v.878a.75.75 0 01-.75.75h-4.5A.75.75 0 015 6.25v-.878zm3.75 7.378a.75.75 0 11-1.5 0 .75.75 0 011.5 0zm3-8.75a.75.75 0 100-1.5.75.75 0 000 1.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-forked%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 21a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zm-3.25-1.75a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0zm-3-12.75a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM2.5 4.75a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0zM18.25 6.5a1.75 1.75 0 110-3.5 1.75 1.75 0 010 3.5zM15 4.75a3.25 3.25 0 106.5 0 3.25 3.25 0 00-6.5 0z\"></path><path fill-rule=\"evenodd\" d=\"M6.5 7.75v1A2.25 2.25 0 008.75 11h6.5a2.25 2.25 0 002.25-2.25v-1H19v1a3.75 3.75 0 01-3.75 3.75h-6.5A3.75 3.75 0 015 8.75v-1h1.5z\"></path><path fill-rule=\"evenodd\" d=\"M11.25 16.25v-5h1.5v5h-1.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// RepoLocked returns a string representing an "repo-locked" Octicon.
func RepoLocked(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-locked%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M1 2.5A2.5 2.5 0 013.5 0h8.75a.75.75 0 01.75.75v3.5a.75.75 0 01-1.5 0V1.5h-8a1 1 0 00-1 1v6.708A2.492 2.492 0 013.5 9h2.75a.75.75 0 010 1.5H3.5a1 1 0 100 2h2.75a.75.75 0 010 1.5H3.5A2.5 2.5 0 011 11.5v-9z\"></path><path fill-rule=\"evenodd\" d=\"M9 10.168V9a3 3 0 116 0v1.168c.591.281 1 .884 1 1.582v2.5A1.75 1.75 0 0114.25 16h-4.5A1.75 1.75 0 018 14.25v-2.5c0-.698.409-1.3 1-1.582zM13.5 10h-3V9a1.5 1.5 0 013 0v1z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// RepoPull returns a string representing an "repo-pull" Octicon.
func RepoPull(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-pull%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13 8V6H7V4h6V2l3 3-3 3zM4 2H3v1h1V2zm7 5h1v6c0 .55-.45 1-1 1H6v2l-1.5-1.5L3 16v-2H1c-.55 0-1-.45-1-1V1c0-.55.45-1 1-1h10c.55 0 1 .45 1 1v2h-1V1H2v9h9V7zm0 4H1v2h2v-1h3v1h5v-2zM4 6H3v1h1V6zm0-2H3v1h1V4zM3 9h1V8H3v1z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// RepoPush returns a string representing an "repo-push" Octicon.
func RepoPush(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-push%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1 2.5A2.5 2.5 0 013.5 0h8.75a.75.75 0 01.75.75v3.5a.75.75 0 01-1.5 0V1.5h-8a1 1 0 00-1 1v6.708A2.492 2.492 0 013.5 9h3.25a.75.75 0 010 1.5H3.5a1 1 0 100 2h5.75a.75.75 0 010 1.5H3.5A2.5 2.5 0 011 11.5v-9zm13.23 7.79a.75.75 0 001.06-1.06l-2.505-2.505a.75.75 0 00-1.06 0L9.22 9.229a.75.75 0 001.06 1.061l1.225-1.224v6.184a.75.75 0 001.5 0V9.066l1.224 1.224z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-push%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M4.75 0A2.75 2.75 0 002 2.75v16.5A2.75 2.75 0 004.75 22h11a.75.75 0 000-1.5h-11c-.69 0-1.25-.56-1.25-1.25V18A1.5 1.5 0 015 16.5h7.25a.75.75 0 000-1.5H5c-.546 0-1.059.146-1.5.401V2.75c0-.69.56-1.25 1.25-1.25H18.5v7a.75.75 0 001.5 0V.75a.75.75 0 00-.75-.75H4.75z\"></path><path d=\"M20 13.903l2.202 2.359a.75.75 0 001.096-1.024l-3.5-3.75a.75.75 0 00-1.096 0l-3.5 3.75a.75.75 0 101.096 1.024l2.202-2.36v9.348a.75.75 0 001.5 0v-9.347z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// RepoTemplate returns a string representing an "repo-template" Octicon.
func RepoTemplate(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-template%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 .75A.75.75 0 016.75 0h2.5a.75.75 0 010 1.5h-2.5A.75.75 0 016 .75zm5 0a.75.75 0 01.75-.75h1.5a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0V1.5h-.75A.75.75 0 0111 .75zM4.992.662a.75.75 0 01-.636.848c-.436.063-.783.41-.846.846a.75.75 0 01-1.485-.212A2.501 2.501 0 014.144.025a.75.75 0 01.848.637zM2.75 4a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 012.75 4zm10.5 0a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5a.75.75 0 01.75-.75zM2.75 8a.75.75 0 01.75.75v.268A1.72 1.72 0 013.75 9h.5a.75.75 0 010 1.5h-.5a.25.25 0 00-.25.25v.75c0 .28.114.532.3.714a.75.75 0 01-1.05 1.072A2.495 2.495 0 012 11.5V8.75A.75.75 0 012.75 8zm10.5 0a.75.75 0 01.75.75v4.5a.75.75 0 01-.75.75h-2.5a.75.75 0 010-1.5h1.75v-2h-.75a.75.75 0 010-1.5h.75v-.25a.75.75 0 01.75-.75zM6 9.75A.75.75 0 016.75 9h2.5a.75.75 0 010 1.5h-2.5A.75.75 0 016 9.75zm-1 2.5v3.25a.25.25 0 00.4.2l1.45-1.087a.25.25 0 01.3 0L8.6 15.7a.25.25 0 00.4-.2v-3.25a.25.25 0 00-.25-.25h-3.5a.25.25 0 00-.25.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-repo-template%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M5.75 0A2.75 2.75 0 003 2.75v1a.75.75 0 001.5 0v-1c0-.69.56-1.25 1.25-1.25h1a.75.75 0 000-1.5h-1zm4 0a.75.75 0 000 1.5h4.5a.75.75 0 000-1.5h-4.5zm7.5 0a.75.75 0 000 1.5h2.25v2.25a.75.75 0 001.5 0v-3a.75.75 0 00-.75-.75h-3zM4.5 6.5a.75.75 0 00-1.5 0v3.75a.75.75 0 001.5 0V6.5zm16.5 0a.75.75 0 00-1.5 0v3.75a.75.75 0 001.5 0V6.5zM4.5 13.25a.75.75 0 00-1.5 0v5.5a3.25 3.25 0 001.95 2.98.75.75 0 10.6-1.375A1.75 1.75 0 014.5 18.75V18A1.5 1.5 0 016 16.5h.75a.75.75 0 000-1.5H6c-.546 0-1.059.146-1.5.401V13.25zm16.5 0a.75.75 0 00-1.5 0V15h-2.25a.75.75 0 000 1.5h2.25v4h-5.25a.75.75 0 000 1.5h6a.75.75 0 00.75-.75v-8zM9.75 15a.75.75 0 000 1.5h4.5a.75.75 0 000-1.5h-4.5zm-2.353 8.461A.25.25 0 017 23.26v-5.01a.25.25 0 01.25-.25h5a.25.25 0 01.25.25v5.01a.25.25 0 01-.397.201l-2.206-1.604a.25.25 0 00-.294 0L7.397 23.46z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Report returns a string representing an "report" Octicon.
func Report(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-report%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 1.5a.25.25 0 00-.25.25v9.5c0 .138.112.25.25.25h2a.75.75 0 01.75.75v2.19l2.72-2.72a.75.75 0 01.53-.22h6.5a.25.25 0 00.25-.25v-9.5a.25.25 0 00-.25-.25H1.75zM0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v9.5A1.75 1.75 0 0114.25 13H8.06l-2.573 2.573A1.457 1.457 0 013 14.543V13H1.75A1.75 1.75 0 010 11.25v-9.5zM9 9a1 1 0 11-2 0 1 1 0 012 0zm-.25-5.25a.75.75 0 00-1.5 0v2.5a.75.75 0 001.5 0v-2.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-report%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.25 4a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25h2.5a.75.75 0 01.75.75v3.19l3.427-3.427A1.75 1.75 0 0111.164 17h9.586a.25.25 0 00.25-.25V4.25a.25.25 0 00-.25-.25H3.25zm-1.75.25c0-.966.784-1.75 1.75-1.75h17.5c.966 0 1.75.784 1.75 1.75v12.5a1.75 1.75 0 01-1.75 1.75h-9.586a.25.25 0 00-.177.073l-3.5 3.5A1.457 1.457 0 015 21.043V18.5H3.25a1.75 1.75 0 01-1.75-1.75V4.25zM12 6a.75.75 0 01.75.75v4a.75.75 0 01-1.5 0v-4A.75.75 0 0112 6zm0 9a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Rocket returns a string representing an "rocket" Octicon.
func Rocket(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-rocket%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M14.064 0a8.75 8.75 0 00-6.187 2.563l-.459.458c-.314.314-.616.641-.904.979H3.31a1.75 1.75 0 00-1.49.833L.11 7.607a.75.75 0 00.418 1.11l3.102.954c.037.051.079.1.124.145l2.429 2.428c.046.046.094.088.145.125l.954 3.102a.75.75 0 001.11.418l2.774-1.707a1.75 1.75 0 00.833-1.49V9.485c.338-.288.665-.59.979-.904l.458-.459A8.75 8.75 0 0016 1.936V1.75A1.75 1.75 0 0014.25 0h-.186zM10.5 10.625c-.088.06-.177.118-.266.175l-2.35 1.521.548 1.783 1.949-1.2a.25.25 0 00.119-.213v-2.066zM3.678 8.116L5.2 5.766c.058-.09.117-.178.176-.266H3.309a.25.25 0 00-.213.119l-1.2 1.95 1.782.547zm5.26-4.493A7.25 7.25 0 0114.063 1.5h.186a.25.25 0 01.25.25v.186a7.25 7.25 0 01-2.123 5.127l-.459.458a15.21 15.21 0 01-2.499 2.02l-2.317 1.5-2.143-2.143 1.5-2.317a15.25 15.25 0 012.02-2.5l.458-.458h.002zM12 5a1 1 0 11-2 0 1 1 0 012 0zm-8.44 9.56a1.5 1.5 0 10-2.12-2.12c-.734.73-1.047 2.332-1.15 3.003a.23.23 0 00.265.265c.671-.103 2.273-.416 3.005-1.148z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-rocket%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M20.322.75a10.75 10.75 0 00-7.373 2.926l-1.304 1.23A23.743 23.743 0 0010.103 6.5H5.066a1.75 1.75 0 00-1.5.85l-2.71 4.514a.75.75 0 00.49 1.12l4.571.963c.039.049.082.096.129.14L8.04 15.96l1.872 1.994c.044.047.091.09.14.129l.963 4.572a.75.75 0 001.12.488l4.514-2.709a1.75 1.75 0 00.85-1.5v-5.038a23.741 23.741 0 001.596-1.542l1.228-1.304a10.75 10.75 0 002.925-7.374V2.499A1.75 1.75 0 0021.498.75h-1.177zM16 15.112c-.333.248-.672.487-1.018.718l-3.393 2.262.678 3.223 3.612-2.167a.25.25 0 00.121-.214v-3.822zm-10.092-2.7L8.17 9.017c.23-.346.47-.685.717-1.017H5.066a.25.25 0 00-.214.121l-2.167 3.612 3.223.679zm8.07-7.644a9.25 9.25 0 016.344-2.518h1.177a.25.25 0 01.25.25v1.176a9.25 9.25 0 01-2.517 6.346l-1.228 1.303a22.248 22.248 0 01-3.854 3.257l-3.288 2.192-1.743-1.858a.764.764 0 00-.034-.034l-1.859-1.744 2.193-3.29a22.248 22.248 0 013.255-3.851l1.304-1.23zM17.5 8a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zm-11 13c.9-.9.9-2.6 0-3.5-.9-.9-2.6-.9-3.5 0-1.209 1.209-1.445 3.901-1.49 4.743a.232.232 0 00.247.247c.842-.045 3.534-.281 4.743-1.49z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Rows returns a string representing an "rows" Octicon.
func Rows(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-rows%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M16 2.75A1.75 1.75 0 0014.25 1H1.75A1.75 1.75 0 000 2.75v2.5A1.75 1.75 0 001.75 7h12.5A1.75 1.75 0 0016 5.25v-2.5zm-1.75-.25a.25.25 0 01.25.25v2.5a.25.25 0 01-.25.25H1.75a.25.25 0 01-.25-.25v-2.5a.25.25 0 01.25-.25h12.5zM16 10.75A1.75 1.75 0 0014.25 9H1.75A1.75 1.75 0 000 10.75v2.5A1.75 1.75 0 001.75 15h12.5A1.75 1.75 0 0016 13.25v-2.5zm-1.75-.25a.25.25 0 01.25.25v2.5a.25.25 0 01-.25.25H1.75a.25.25 0 01-.25-.25v-2.5a.25.25 0 01.25-.25h12.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-rows%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M22 3.75A1.75 1.75 0 0020.25 2H3.75A1.75 1.75 0 002 3.75v5.5c0 .966.784 1.75 1.75 1.75h16.5A1.75 1.75 0 0022 9.25v-5.5zm-1.75-.25a.25.25 0 01.25.25v5.5a.25.25 0 01-.25.25H3.75a.25.25 0 01-.25-.25v-5.5a.25.25 0 01.25-.25h16.5zM22 14.75A1.75 1.75 0 0020.25 13H3.75A1.75 1.75 0 002 14.75v5.5c0 .966.784 1.75 1.75 1.75h16.5A1.75 1.75 0 0022 20.25v-5.5zm-1.75-.25a.25.25 0 01.25.25v5.5a.25.25 0 01-.25.25H3.75a.25.25 0 01-.25-.25v-5.5a.25.25 0 01.25-.25h16.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Rss returns a string representing an "rss" Octicon.
func Rss(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-rss%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.002 2.725a.75.75 0 01.797-.699C8.79 2.42 13.58 7.21 13.974 13.201a.75.75 0 11-1.497.098 10.502 10.502 0 00-9.776-9.776.75.75 0 01-.7-.798zM2 13a1 1 0 112 0 1 1 0 01-2 0zm.84-5.95a.75.75 0 00-.179 1.489c2.509.3 4.5 2.291 4.8 4.8a.75.75 0 101.49-.178A7.003 7.003 0 002.838 7.05z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-rss%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.5 3.25a.75.75 0 01.75-.75C14.053 2.5 22 10.447 22 20.25a.75.75 0 01-1.5 0C20.5 11.275 13.225 4 4.25 4a.75.75 0 01-.75-.75zM3.5 19a2 2 0 114 0 2 2 0 01-4 0zm.75-9.5a.75.75 0 000 1.5 9.25 9.25 0 019.25 9.25.75.75 0 001.5 0C15 14.313 10.187 9.5 4.25 9.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Ruby returns a string representing an "ruby" Octicon.
func Ruby(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-ruby%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.637 2.291A.75.75 0 014.23 2h7.54a.75.75 0 01.593.291l3.48 4.5a.75.75 0 01-.072.999l-7.25 7a.75.75 0 01-1.042 0l-7.25-7a.75.75 0 01-.072-.999l3.48-4.5zM4.598 3.5L1.754 7.177 8 13.207l6.246-6.03L11.402 3.5H4.598z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-ruby%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.873 3.26A.75.75 0 016.44 3h11.31a.75.75 0 01.576.27l5 6a.75.75 0 01-.028.992l-10.75 11.5a.75.75 0 01-1.096 0l-10.75-11.5a.75.75 0 01-.02-1.003l5.19-6zm.91 1.24L2.258 9.73 12 20.153l9.75-10.43L17.399 4.5H6.783z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ScreenFull returns a string representing an "screen-full" Octicon.
func ScreenFull(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-screen-full%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.75 2.5a.25.25 0 00-.25.25v2.5a.75.75 0 01-1.5 0v-2.5C1 1.784 1.784 1 2.75 1h2.5a.75.75 0 010 1.5h-2.5zM10 1.75a.75.75 0 01.75-.75h2.5c.966 0 1.75.784 1.75 1.75v2.5a.75.75 0 01-1.5 0v-2.5a.25.25 0 00-.25-.25h-2.5a.75.75 0 01-.75-.75zM1.75 10a.75.75 0 01.75.75v2.5c0 .138.112.25.25.25h2.5a.75.75 0 010 1.5h-2.5A1.75 1.75 0 011 13.25v-2.5a.75.75 0 01.75-.75zm12.5 0a.75.75 0 01.75.75v2.5A1.75 1.75 0 0113.25 15h-2.5a.75.75 0 010-1.5h2.5a.25.25 0 00.25-.25v-2.5a.75.75 0 01.75-.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-screen-full%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 4.5a.25.25 0 00-.25.25v3.5a.75.75 0 01-1.5 0v-3.5C3 3.784 3.784 3 4.75 3h3.5a.75.75 0 010 1.5h-3.5zM15 3.75a.75.75 0 01.75-.75h3.5c.966 0 1.75.784 1.75 1.75v3.5a.75.75 0 01-1.5 0v-3.5a.25.25 0 00-.25-.25h-3.5a.75.75 0 01-.75-.75zM3.75 15a.75.75 0 01.75.75v3.5c0 .138.112.25.25.25h3.5a.75.75 0 010 1.5h-3.5A1.75 1.75 0 013 19.25v-3.5a.75.75 0 01.75-.75zm16.5 0a.75.75 0 01.75.75v3.5A1.75 1.75 0 0119.25 21h-3.5a.75.75 0 010-1.5h3.5a.25.25 0 00.25-.25v-3.5a.75.75 0 01.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ScreenNormal returns a string representing an "screen-normal" Octicon.
func ScreenNormal(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-screen-normal%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.25 1a.75.75 0 01.75.75v2.5A1.75 1.75 0 014.25 6h-2.5a.75.75 0 010-1.5h2.5a.25.25 0 00.25-.25v-2.5A.75.75 0 015.25 1zm5.5 0a.75.75 0 01.75.75v2.5c0 .138.112.25.25.25h2.5a.75.75 0 010 1.5h-2.5A1.75 1.75 0 0110 4.25v-2.5a.75.75 0 01.75-.75zM1 10.75a.75.75 0 01.75-.75h2.5c.966 0 1.75.784 1.75 1.75v2.5a.75.75 0 01-1.5 0v-2.5a.25.25 0 00-.25-.25h-2.5a.75.75 0 01-.75-.75zm9 1c0-.966.784-1.75 1.75-1.75h2.5a.75.75 0 010 1.5h-2.5a.25.25 0 00-.25.25v2.5a.75.75 0 01-1.5 0v-2.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-screen-normal%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.25 3a.75.75 0 01.75.75v3.5A1.75 1.75 0 017.25 9h-3.5a.75.75 0 010-1.5h3.5a.25.25 0 00.25-.25v-3.5A.75.75 0 018.25 3zm7.5 0a.75.75 0 01.75.75v3.5c0 .138.112.25.25.25h3.5a.75.75 0 010 1.5h-3.5A1.75 1.75 0 0115 7.25v-3.5a.75.75 0 01.75-.75zM3 15.75a.75.75 0 01.75-.75h3.5c.966 0 1.75.784 1.75 1.75v3.5a.75.75 0 01-1.5 0v-3.5a.25.25 0 00-.25-.25h-3.5a.75.75 0 01-.75-.75zm12 1c0-.966.784-1.75 1.75-1.75h3.5a.75.75 0 010 1.5h-3.5a.25.25 0 00-.25.25v3.5a.75.75 0 01-1.5 0v-3.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Search returns a string representing an "search" Octicon.
func Search(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-search%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.5 7a4.499 4.499 0 11-8.998 0A4.499 4.499 0 0111.5 7zm-.82 4.74a6 6 0 111.06-1.06l3.04 3.04a.75.75 0 11-1.06 1.06l-3.04-3.04z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-search%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M14.53 15.59a8.25 8.25 0 111.06-1.06l5.69 5.69a.75.75 0 11-1.06 1.06l-5.69-5.69zM2.5 9.25a6.75 6.75 0 1111.74 4.547.746.746 0 00-.443.442A6.75 6.75 0 012.5 9.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Server returns a string representing an "server" Octicon.
func Server(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-server%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 1A1.75 1.75 0 000 2.75v4c0 .372.116.717.314 1a1.742 1.742 0 00-.314 1v4c0 .966.784 1.75 1.75 1.75h12.5A1.75 1.75 0 0016 12.75v-4c0-.372-.116-.717-.314-1 .198-.283.314-.628.314-1v-4A1.75 1.75 0 0014.25 1H1.75zm0 7.5a.25.25 0 00-.25.25v4c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25v-4a.25.25 0 00-.25-.25H1.75zM1.5 2.75a.25.25 0 01.25-.25h12.5a.25.25 0 01.25.25v4a.25.25 0 01-.25.25H1.75a.25.25 0 01-.25-.25v-4zm5.5 2A.75.75 0 017.75 4h4.5a.75.75 0 010 1.5h-4.5A.75.75 0 017 4.75zM7.75 10a.75.75 0 000 1.5h4.5a.75.75 0 000-1.5h-4.5zM3 4.75A.75.75 0 013.75 4h.5a.75.75 0 010 1.5h-.5A.75.75 0 013 4.75zM3.75 10a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-server%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M10.75 6.5a.75.75 0 000 1.5h6.5a.75.75 0 000-1.5h-6.5zM6 7.25a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5A.75.75 0 016 7.25zm4 9a.75.75 0 01.75-.75h6.5a.75.75 0 010 1.5h-6.5a.75.75 0 01-.75-.75zm-3.25-.75a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5z\"></path><path fill-rule=\"evenodd\" d=\"M3.25 2A1.75 1.75 0 001.5 3.75v7c0 .372.116.716.314 1a1.742 1.742 0 00-.314 1v7c0 .966.784 1.75 1.75 1.75h17.5a1.75 1.75 0 001.75-1.75v-7c0-.372-.116-.716-.314-1 .198-.284.314-.628.314-1v-7A1.75 1.75 0 0020.75 2H3.25zm0 9h17.5a.25.25 0 00.25-.25v-7a.25.25 0 00-.25-.25H3.25a.25.25 0 00-.25.25v7c0 .138.112.25.25.25zm0 1.5a.25.25 0 00-.25.25v7c0 .138.112.25.25.25h17.5a.25.25 0 00.25-.25v-7a.25.25 0 00-.25-.25H3.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Share returns a string representing an "share" Octicon.
func Share(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-share%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.823.177L4.927 3.073a.25.25 0 00.177.427H7.25v5.75a.75.75 0 001.5 0V3.5h2.146a.25.25 0 00.177-.427L8.177.177a.25.25 0 00-.354 0zM3.75 6.5a.25.25 0 00-.25.25v6.5c0 .138.112.25.25.25h8.5a.25.25 0 00.25-.25v-6.5a.25.25 0 00-.25-.25h-1a.75.75 0 010-1.5h1c.966 0 1.75.784 1.75 1.75v6.5A1.75 1.75 0 0112.25 15h-8.5A1.75 1.75 0 012 13.25v-6.5C2 5.784 2.784 5 3.75 5h1a.75.75 0 110 1.5h-1z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-share%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.53 1.22a.75.75 0 00-1.06 0L8.22 4.47a.75.75 0 001.06 1.06l1.97-1.97v10.69a.75.75 0 001.5 0V3.56l1.97 1.97a.75.75 0 101.06-1.06l-3.25-3.25zM5.5 9.75a.25.25 0 01.25-.25h2.5a.75.75 0 000-1.5h-2.5A1.75 1.75 0 004 9.75v10.5c0 .966.784 1.75 1.75 1.75h12.5A1.75 1.75 0 0020 20.25V9.75A1.75 1.75 0 0018.25 8h-2.5a.75.75 0 000 1.5h2.5a.25.25 0 01.25.25v10.5a.25.25 0 01-.25.25H5.75a.25.25 0 01-.25-.25V9.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ShareAndroid returns a string representing an "share-android" Octicon.
func ShareAndroid(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-share-android%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M13.5 3a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM15 3a3 3 0 01-5.175 2.066l-3.92 2.179a3.005 3.005 0 010 1.51l3.92 2.179a3 3 0 11-.73 1.31l-3.92-2.178a3 3 0 110-4.133l3.92-2.178A3 3 0 1115 3zm-1.5 10a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zm-9-5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-share-android%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M20 5.5a3.5 3.5 0 01-6.062 2.385l-5.112 3.021a3.497 3.497 0 010 2.188l5.112 3.021a3.5 3.5 0 11-.764 1.29l-5.112-3.02a3.5 3.5 0 110-4.77l5.112-3.021v.001A3.5 3.5 0 1120 5.5zm-1.5 0a2 2 0 11-4 0 2 2 0 014 0zM5.5 14a2 2 0 100-4 2 2 0 000 4zm13 4.5a2 2 0 11-4 0 2 2 0 014 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Shield returns a string representing an "shield" Octicon.
func Shield(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.467.133a1.75 1.75 0 011.066 0l5.25 1.68A1.75 1.75 0 0115 3.48V7c0 1.566-.32 3.182-1.303 4.682-.983 1.498-2.585 2.813-5.032 3.855a1.7 1.7 0 01-1.33 0c-2.447-1.042-4.049-2.357-5.032-3.855C1.32 10.182 1 8.566 1 7V3.48a1.75 1.75 0 011.217-1.667l5.25-1.68zm.61 1.429a.25.25 0 00-.153 0l-5.25 1.68a.25.25 0 00-.174.238V7c0 1.358.275 2.666 1.057 3.86.784 1.194 2.121 2.34 4.366 3.297a.2.2 0 00.154 0c2.245-.956 3.582-2.104 4.366-3.298C13.225 9.666 13.5 8.36 13.5 7V3.48a.25.25 0 00-.174-.237l-5.25-1.68zM9 10.5a1 1 0 11-2 0 1 1 0 012 0zm-.25-5.75a.75.75 0 10-1.5 0v3a.75.75 0 001.5 0v-3z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M13 15.5a1 1 0 11-2 0 1 1 0 012 0zm-.25-8.25a.75.75 0 00-1.5 0v4.5a.75.75 0 001.5 0v-4.5z\"></path><path fill-rule=\"evenodd\" d=\"M11.46.637a1.75 1.75 0 011.08 0l8.25 2.675A1.75 1.75 0 0122 4.976V10c0 6.19-3.77 10.705-9.401 12.83a1.699 1.699 0 01-1.198 0C5.771 20.704 2 16.19 2 10V4.976c0-.76.49-1.43 1.21-1.664L11.46.637zm.617 1.426a.25.25 0 00-.154 0L3.673 4.74a.249.249 0 00-.173.237V10c0 5.461 3.28 9.483 8.43 11.426a.2.2 0 00.14 0C17.22 19.483 20.5 15.46 20.5 10V4.976a.25.25 0 00-.173-.237l-8.25-2.676z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ShieldCheck returns a string representing an "shield-check" Octicon.
func ShieldCheck(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield-check%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.533.133a1.75 1.75 0 00-1.066 0l-5.25 1.68A1.75 1.75 0 001 3.48V7c0 1.566.32 3.182 1.303 4.682.983 1.498 2.585 2.813 5.032 3.855a1.7 1.7 0 001.33 0c2.447-1.042 4.049-2.357 5.032-3.855C14.68 10.182 15 8.566 15 7V3.48a1.75 1.75 0 00-1.217-1.667L8.533.133zm-.61 1.429a.25.25 0 01.153 0l5.25 1.68a.25.25 0 01.174.238V7c0 1.358-.275 2.666-1.057 3.86-.784 1.194-2.121 2.34-4.366 3.297a.2.2 0 01-.154 0c-2.245-.956-3.582-2.104-4.366-3.298C2.775 9.666 2.5 8.36 2.5 7V3.48a.25.25 0 01.174-.237l5.25-1.68zM11.28 6.28a.75.75 0 00-1.06-1.06L7.25 8.19l-.97-.97a.75.75 0 10-1.06 1.06l1.5 1.5a.75.75 0 001.06 0l3.5-3.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield-check%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M16.53 9.78a.75.75 0 00-1.06-1.06L11 13.19l-1.97-1.97a.75.75 0 00-1.06 1.06l2.5 2.5a.75.75 0 001.06 0l5-5z\"></path><path fill-rule=\"evenodd\" d=\"M12.54.637a1.75 1.75 0 00-1.08 0L3.21 3.312A1.75 1.75 0 002 4.976V10c0 6.19 3.77 10.705 9.401 12.83.386.145.812.145 1.198 0C18.229 20.704 22 16.19 22 10V4.976c0-.759-.49-1.43-1.21-1.664L12.54.637zm-.617 1.426a.25.25 0 01.154 0l8.25 2.676a.25.25 0 01.173.237V10c0 5.461-3.28 9.483-8.43 11.426a.2.2 0 01-.14 0C6.78 19.483 3.5 15.46 3.5 10V4.976c0-.108.069-.203.173-.237l8.25-2.676z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ShieldLock returns a string representing an "shield-lock" Octicon.
func ShieldLock(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield-lock%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.533.133a1.75 1.75 0 00-1.066 0l-5.25 1.68A1.75 1.75 0 001 3.48V7c0 1.566.32 3.182 1.303 4.682.983 1.498 2.585 2.813 5.032 3.855a1.7 1.7 0 001.33 0c2.447-1.042 4.049-2.357 5.032-3.855C14.68 10.182 15 8.566 15 7V3.48a1.75 1.75 0 00-1.217-1.667L8.533.133zm-.61 1.429a.25.25 0 01.153 0l5.25 1.68a.25.25 0 01.174.238V7c0 1.358-.275 2.666-1.057 3.86-.784 1.194-2.121 2.34-4.366 3.297a.2.2 0 01-.154 0c-2.245-.956-3.582-2.104-4.366-3.298C2.775 9.666 2.5 8.36 2.5 7V3.48a.25.25 0 01.174-.237l5.25-1.68zM9.5 6.5a1.5 1.5 0 01-.75 1.3v2.45a.75.75 0 01-1.5 0V7.8A1.5 1.5 0 119.5 6.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield-lock%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.077 2.563a.25.25 0 00-.154 0L3.673 5.24a.249.249 0 00-.173.237V10.5c0 5.461 3.28 9.483 8.43 11.426a.2.2 0 00.14 0c5.15-1.943 8.43-5.965 8.43-11.426V5.476a.25.25 0 00-.173-.237l-8.25-2.676zm-.617-1.426a1.75 1.75 0 011.08 0l8.25 2.675A1.75 1.75 0 0122 5.476V10.5c0 6.19-3.77 10.705-9.401 12.83a1.699 1.699 0 01-1.198 0C5.771 21.204 2 16.69 2 10.5V5.476c0-.76.49-1.43 1.21-1.664l8.25-2.675zM13 12.232A2 2 0 0012 8.5a2 2 0 00-1 3.732V15a1 1 0 102 0v-2.768z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ShieldX returns a string representing an "shield-x" Octicon.
func ShieldX(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield-x%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.533.133a1.75 1.75 0 00-1.066 0l-5.25 1.68A1.75 1.75 0 001 3.48V7c0 1.566.32 3.182 1.303 4.682.983 1.498 2.585 2.813 5.032 3.855a1.7 1.7 0 001.33 0c2.447-1.042 4.049-2.357 5.032-3.855C14.68 10.182 15 8.566 15 7V3.48a1.75 1.75 0 00-1.217-1.667L8.533.133zm-.61 1.429a.25.25 0 01.153 0l5.25 1.68a.25.25 0 01.174.238V7c0 1.358-.275 2.666-1.057 3.86-.784 1.194-2.121 2.34-4.366 3.297a.2.2 0 01-.154 0c-2.245-.956-3.582-2.104-4.366-3.298C2.775 9.666 2.5 8.36 2.5 7V3.48a.25.25 0 01.174-.237l5.25-1.68zM6.78 5.22a.75.75 0 10-1.06 1.06L6.94 7.5 5.72 8.72a.75.75 0 001.06 1.06L8 8.56l1.22 1.22a.75.75 0 101.06-1.06L9.06 7.5l1.22-1.22a.75.75 0 10-1.06-1.06L8 6.44 6.78 5.22z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-shield-x%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M9.28 7.72a.75.75 0 00-1.06 1.06l2.72 2.72-2.72 2.72a.75.75 0 101.06 1.06L12 12.56l2.72 2.72a.75.75 0 101.06-1.06l-2.72-2.72 2.72-2.72a.75.75 0 00-1.06-1.06L12 10.44 9.28 7.72z\"></path><path fill-rule=\"evenodd\" d=\"M12.54.637a1.75 1.75 0 00-1.08 0L3.21 3.312A1.75 1.75 0 002 4.976V10c0 6.19 3.77 10.705 9.401 12.83.386.145.812.145 1.198 0C18.229 20.704 22 16.19 22 10V4.976c0-.759-.49-1.43-1.21-1.664L12.54.637zm-.617 1.426a.25.25 0 01.154 0l8.25 2.676a.25.25 0 01.173.237V10c0 5.461-3.28 9.483-8.43 11.426a.2.2 0 01-.14 0C6.78 19.483 3.5 15.46 3.5 10V4.976c0-.108.069-.203.173-.237l8.25-2.676z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SidebarCollapse returns a string representing an "sidebar-collapse" Octicon.
func SidebarCollapse(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sidebar-collapse%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.823 7.823L4.427 5.427A.25.25 0 004 5.604v4.792c0 .223.27.335.427.177l2.396-2.396a.25.25 0 000-.354z\"></path><path fill-rule=\"evenodd\" d=\"M1.75 0A1.75 1.75 0 000 1.75v12.5C0 15.216.784 16 1.75 16h12.5A1.75 1.75 0 0016 14.25V1.75A1.75 1.75 0 0014.25 0H1.75zM1.5 1.75a.25.25 0 01.25-.25H9.5v13H1.75a.25.25 0 01-.25-.25V1.75zM11 14.5v-13h3.25a.25.25 0 01.25.25v12.5a.25.25 0 01-.25.25H11z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sidebar-collapse%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.22 14.47L9.69 12 7.22 9.53a.75.75 0 111.06-1.06l3 3a.75.75 0 010 1.06l-3 3a.75.75 0 01-1.06-1.06z\"></path><path fill-rule=\"evenodd\" d=\"M3.75 2A1.75 1.75 0 002 3.75v16.5c0 .966.784 1.75 1.75 1.75h16.5A1.75 1.75 0 0022 20.25V3.75A1.75 1.75 0 0020.25 2H3.75zM3.5 3.75a.25.25 0 01.25-.25H15v17H3.75a.25.25 0 01-.25-.25V3.75zm13 16.75v-17h3.75a.25.25 0 01.25.25v16.5a.25.25 0 01-.25.25H16.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SidebarExpand returns a string representing an "sidebar-expand" Octicon.
func SidebarExpand(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sidebar-expand%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.177 7.823l2.396-2.396A.25.25 0 017 5.604v4.792a.25.25 0 01-.427.177L4.177 8.177a.25.25 0 010-.354z\"></path><path fill-rule=\"evenodd\" d=\"M0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v12.5A1.75 1.75 0 0114.25 16H1.75A1.75 1.75 0 010 14.25V1.75zm1.75-.25a.25.25 0 00-.25.25v12.5c0 .138.112.25.25.25H9.5v-13H1.75zm12.5 13H11v-13h3.25a.25.25 0 01.25.25v12.5a.25.25 0 01-.25.25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sidebar-expand%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.28 9.53L8.81 12l2.47 2.47a.75.75 0 11-1.06 1.06l-3-3a.75.75 0 010-1.06l3-3a.75.75 0 111.06 1.06z\"></path><path fill-rule=\"evenodd\" d=\"M3.75 2A1.75 1.75 0 002 3.75v16.5c0 .966.784 1.75 1.75 1.75h16.5A1.75 1.75 0 0022 20.25V3.75A1.75 1.75 0 0020.25 2H3.75zM3.5 3.75a.25.25 0 01.25-.25H15v17H3.75a.25.25 0 01-.25-.25V3.75zm13 16.75v-17h3.75a.25.25 0 01.25.25v16.5a.25.25 0 01-.25.25H16.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SignIn returns a string representing an "sign-in" Octicon.
func SignIn(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sign-in%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 2.75C2 1.784 2.784 1 3.75 1h2.5a.75.75 0 010 1.5h-2.5a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h2.5a.75.75 0 010 1.5h-2.5A1.75 1.75 0 012 13.25V2.75zm6.56 4.5l1.97-1.97a.75.75 0 10-1.06-1.06L6.22 7.47a.75.75 0 000 1.06l3.25 3.25a.75.75 0 101.06-1.06L8.56 8.75h5.69a.75.75 0 000-1.5H8.56z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sign-in%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3 3.25c0-.966.784-1.75 1.75-1.75h5.5a.75.75 0 010 1.5h-5.5a.25.25 0 00-.25.25v17.5c0 .138.112.25.25.25h5.5a.75.75 0 010 1.5h-5.5A1.75 1.75 0 013 20.75V3.25zm9.994 9.5l3.3 3.484a.75.75 0 01-1.088 1.032l-4.5-4.75a.75.75 0 010-1.032l4.5-4.75a.75.75 0 011.088 1.032l-3.3 3.484h8.256a.75.75 0 010 1.5h-8.256z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SignOut returns a string representing an "sign-out" Octicon.
func SignOut(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sign-out%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 2.75C2 1.784 2.784 1 3.75 1h2.5a.75.75 0 010 1.5h-2.5a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h2.5a.75.75 0 010 1.5h-2.5A1.75 1.75 0 012 13.25V2.75zm10.44 4.5H6.75a.75.75 0 000 1.5h5.69l-1.97 1.97a.75.75 0 101.06 1.06l3.25-3.25a.75.75 0 000-1.06l-3.25-3.25a.75.75 0 10-1.06 1.06l1.97 1.97z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sign-out%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3 3.25c0-.966.784-1.75 1.75-1.75h5.5a.75.75 0 010 1.5h-5.5a.25.25 0 00-.25.25v17.5c0 .138.112.25.25.25h5.5a.75.75 0 010 1.5h-5.5A1.75 1.75 0 013 20.75V3.25zm16.006 9.5l-3.3 3.484a.75.75 0 001.088 1.032l4.5-4.75a.75.75 0 000-1.032l-4.5-4.75a.75.75 0 00-1.088 1.032l3.3 3.484H10.75a.75.75 0 000 1.5h8.256z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SingleSelect returns a string representing an "single-select" Octicon.
func SingleSelect(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-single-select%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M5.06 7.356l2.795 2.833c.08.081.21.081.29 0l2.794-2.833c.13-.131.038-.356-.145-.356H5.206c-.183 0-.275.225-.145.356z\"></path><path fill-rule=\"evenodd\" d=\"M1 2.75C1 1.784 1.784 1 2.75 1h10.5c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0113.25 15H2.75A1.75 1.75 0 011 13.25V2.75zm1.75-.25a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h10.5a.25.25 0 00.25-.25V2.75a.25.25 0 00-.25-.25H2.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-single-select%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M7.854 10.854l3.792 3.792a.5.5 0 00.708 0l3.793-3.792a.5.5 0 00-.354-.854H8.207a.5.5 0 00-.353.854z\"></path><path fill-rule=\"evenodd\" d=\"M2 3.75C2 2.784 2.784 2 3.75 2h16.5c.966 0 1.75.784 1.75 1.75v16.5A1.75 1.75 0 0120.25 22H3.75A1.75 1.75 0 012 20.25V3.75zm1.75-.25a.25.25 0 00-.25.25v16.5c0 .138.112.25.25.25h16.5a.25.25 0 00.25-.25V3.75a.25.25 0 00-.25-.25H3.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Skip returns a string representing an "skip" Octicon.
func Skip(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-skip%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 8a6.5 6.5 0 1113 0 6.5 6.5 0 01-13 0zM8 0a8 8 0 100 16A8 8 0 008 0zm3.28 5.78a.75.75 0 00-1.06-1.06l-5.5 5.5a.75.75 0 101.06 1.06l5.5-5.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-skip%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M17.28 7.78a.75.75 0 00-1.06-1.06l-9.5 9.5a.75.75 0 101.06 1.06l9.5-9.5z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Sliders returns a string representing an "sliders" Octicon.
func Sliders(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sliders%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M15 2.75a.75.75 0 01-.75.75h-4a.75.75 0 010-1.5h4a.75.75 0 01.75.75zm-8.5.75v1.25a.75.75 0 001.5 0v-4a.75.75 0 00-1.5 0V2H1.75a.75.75 0 000 1.5H6.5zm1.25 5.25a.75.75 0 000-1.5h-6a.75.75 0 000 1.5h6zM15 8a.75.75 0 01-.75.75H11.5V10a.75.75 0 11-1.5 0V6a.75.75 0 011.5 0v1.25h2.75A.75.75 0 0115 8zm-9 5.25v-2a.75.75 0 00-1.5 0v1.25H1.75a.75.75 0 000 1.5H4.5v1.25a.75.75 0 001.5 0v-2zm9 0a.75.75 0 01-.75.75h-6a.75.75 0 010-1.5h6a.75.75 0 01.75.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Smiley returns a string representing an "smiley" Octicon.
func Smiley(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-smiley%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.5 8a6.5 6.5 0 1113 0 6.5 6.5 0 01-13 0zM8 0a8 8 0 100 16A8 8 0 008 0zM5 8a1 1 0 100-2 1 1 0 000 2zm7-1a1 1 0 11-2 0 1 1 0 012 0zM5.32 9.636a.75.75 0 011.038.175l.007.009c.103.118.22.222.35.31.264.178.683.37 1.285.37.602 0 1.02-.192 1.285-.371.13-.088.247-.192.35-.31l.007-.008a.75.75 0 111.222.87l-.614-.431c.614.43.614.431.613.431v.001l-.001.002-.002.003-.005.007-.014.019a1.984 1.984 0 01-.184.213c-.16.166-.338.316-.53.445-.63.418-1.37.638-2.127.629-.946 0-1.652-.308-2.126-.63a3.32 3.32 0 01-.715-.657l-.014-.02-.005-.006-.002-.003v-.002h-.001l.613-.432-.614.43a.75.75 0 01.183-1.044h.001z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-smiley%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M8.456 14.494a.75.75 0 011.068.17 3.08 3.08 0 00.572.492A3.381 3.381 0 0012 15.72c.855 0 1.487-.283 1.904-.562a3.081 3.081 0 00.572-.492l.021-.026a.75.75 0 011.197.905l-.027.034c-.013.016-.03.038-.052.063-.044.05-.105.119-.184.198a4.569 4.569 0 01-.695.566A4.88 4.88 0 0112 17.22a4.88 4.88 0 01-2.736-.814 4.57 4.57 0 01-.695-.566 3.253 3.253 0 01-.236-.261c-.259-.332-.223-.824.123-1.084z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path><path d=\"M9 10.75a1.25 1.25 0 11-2.5 0 1.25 1.25 0 012.5 0zM16.25 12a1.25 1.25 0 100-2.5 1.25 1.25 0 000 2.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SortAsc returns a string representing an "sort-asc" Octicon.
func SortAsc(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sort-asc%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 4.25a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5H.75A.75.75 0 010 4.25zm0 4a.75.75 0 01.75-.75h4.5a.75.75 0 010 1.5H.75A.75.75 0 010 8.25zm0 4a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5H.75a.75.75 0 01-.75-.75zm12.927-9.677a.25.25 0 00-.354 0l-3 3A.25.25 0 009.75 6H12v6.75a.75.75 0 001.5 0V6h2.25a.25.25 0 00.177-.427l-3-3z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sort-asc%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M18.5 17.25a.75.75 0 01-1.5 0V7.56l-2.22 2.22a.75.75 0 11-1.06-1.06l3.5-3.5a.75.75 0 011.06 0l3.5 3.5a.75.75 0 01-1.06 1.06L18.5 7.56v9.69zm-15.75.25a.75.75 0 010-1.5h9.5a.75.75 0 010 1.5h-9.5zm0-5a.75.75 0 010-1.5h5.5a.75.75 0 010 1.5h-5.5zm0-5a.75.75 0 010-1.5h3.5a.75.75 0 010 1.5h-3.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SortDesc returns a string representing an "sort-desc" Octicon.
func SortDesc(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sort-desc%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 4.25a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5H.75A.75.75 0 010 4.25zm0 4a.75.75 0 01.75-.75h4.5a.75.75 0 010 1.5H.75A.75.75 0 010 8.25zm0 4a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5H.75a.75.75 0 01-.75-.75z\"></path><path d=\"M13.5 10h2.25a.25.25 0 01.177.427l-3 3a.25.25 0 01-.354 0l-3-3A.25.25 0 019.75 10H12V3.75a.75.75 0 011.5 0V10z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sort-desc%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M18.5 16.44V6.75a.75.75 0 00-1.5 0v9.69l-2.22-2.22a.75.75 0 10-1.06 1.06l3.5 3.5a.75.75 0 001.06 0l3.5-3.5a.75.75 0 10-1.06-1.06l-2.22 2.22zM2 7.25a.75.75 0 01.75-.75h9.5a.75.75 0 010 1.5h-9.5A.75.75 0 012 7.25zm0 5a.75.75 0 01.75-.75h5.5a.75.75 0 010 1.5h-5.5a.75.75 0 01-.75-.75zm0 5a.75.75 0 01.75-.75h3.5a.75.75 0 010 1.5h-3.5a.75.75 0 01-.75-.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Square returns a string representing an "square" Octicon.
func Square(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-square%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4 5.75C4 4.784 4.784 4 5.75 4h4.5c.966 0 1.75.784 1.75 1.75v4.5A1.75 1.75 0 0110.25 12h-4.5A1.75 1.75 0 014 10.25v-4.5zm1.75-.25a.25.25 0 00-.25.25v4.5c0 .138.112.25.25.25h4.5a.25.25 0 00.25-.25v-4.5a.25.25 0 00-.25-.25h-4.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-square%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6 7.75C6 6.784 6.784 6 7.75 6h8.5c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0116.25 18h-8.5A1.75 1.75 0 016 16.25v-8.5zm1.75-.25a.25.25 0 00-.25.25v8.5c0 .138.112.25.25.25h8.5a.25.25 0 00.25-.25v-8.5a.25.25 0 00-.25-.25h-8.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// SquareFill returns a string representing an "square-fill" Octicon.
func SquareFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-square-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.75 4A1.75 1.75 0 004 5.75v4.5c0 .966.784 1.75 1.75 1.75h4.5A1.75 1.75 0 0012 10.25v-4.5A1.75 1.75 0 0010.25 4h-4.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-square-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.75 6A1.75 1.75 0 006 7.75v8.5c0 .966.784 1.75 1.75 1.75h8.5A1.75 1.75 0 0018 16.25v-8.5A1.75 1.75 0 0016.25 6h-8.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Squirrel returns a string representing an "squirrel" Octicon.
func Squirrel(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-squirrel%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.499.75a.75.75 0 011.5 0v.996C5.9 2.903 6.793 3.65 7.662 4.376l.24.202c-.036-.694.055-1.422.426-2.163C9.1.873 10.794-.045 12.622.26 14.408.558 16 1.94 16 4.25c0 1.278-.954 2.575-2.44 2.734l.146.508.065.22c.203.701.412 1.455.476 2.226.142 1.707-.4 3.03-1.487 3.898C11.714 14.671 10.27 15 8.75 15h-6a.75.75 0 010-1.5h1.376a4.489 4.489 0 01-.563-1.191 3.833 3.833 0 01-.05-2.063 4.636 4.636 0 01-2.025-.293.75.75 0 11.525-1.406c1.357.507 2.376-.006 2.698-.318l.009-.01a.748.748 0 011.06 0 .75.75 0 01-.012 1.074c-.912.92-.992 1.835-.768 2.586.221.74.745 1.337 1.196 1.621H8.75c1.343 0 2.398-.296 3.074-.836.635-.507 1.036-1.31.928-2.602-.05-.603-.216-1.224-.422-1.93l-.064-.221c-.12-.407-.246-.84-.353-1.29a2.404 2.404 0 01-.507-.441 3.063 3.063 0 01-.633-1.248.75.75 0 011.455-.364c.046.185.144.436.31.627.146.168.353.305.712.305.738 0 1.25-.615 1.25-1.25 0-1.47-.95-2.315-2.123-2.51-1.172-.196-2.227.387-2.706 1.345-.46.92-.27 1.774.019 3.062l.042.19a.753.753 0 01.01.05c.348.443.666.949.94 1.553a.75.75 0 11-1.365.62c-.553-1.217-1.32-1.94-2.3-2.768a85.08 85.08 0 00-.317-.265c-.814-.68-1.75-1.462-2.692-2.619a3.74 3.74 0 00-1.023.88c-.406.495-.663 1.036-.722 1.508.116.122.306.21.591.239.388.038.797-.06 1.032-.19a.75.75 0 01.728 1.31c-.515.287-1.23.439-1.906.373-.682-.067-1.473-.38-1.879-1.193L.75 5.677V5.5c0-.984.48-1.94 1.077-2.664.46-.559 1.05-1.055 1.673-1.353V.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-squirrel%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M18.377 3.49c-1.862-.31-3.718.62-4.456 2.095-.428.857-.691 1.624-.728 2.361-.035.71.138 1.444.67 2.252.644.854 1.199 1.913 1.608 3.346a.75.75 0 11-1.442.412c-.353-1.236-.82-2.135-1.372-2.865l-.008-.01c-.53-.698-1.14-1.242-1.807-1.778a50.724 50.724 0 00-.667-.524C9.024 7.884 7.71 6.863 6.471 5.16c-.59.287-1.248.798-1.806 1.454-.665.78-1.097 1.66-1.158 2.446.246.36.685.61 1.246.715.643.12 1.278.015 1.633-.182a.75.75 0 11.728 1.311c-.723.402-1.728.516-2.637.346-.916-.172-1.898-.667-2.398-1.666L2 9.427V9.25c0-1.323.678-2.615 1.523-3.607.7-.824 1.59-1.528 2.477-1.917V2.75a.75.75 0 111.5 0v1.27c1.154 1.67 2.363 2.612 3.568 3.551.207.162.415.323.621.489.001-.063.003-.126.006-.188.052-1.034.414-2.017.884-2.958 1.06-2.118 3.594-3.313 6.044-2.904 1.225.204 2.329.795 3.125 1.748C22.546 4.713 23 5.988 23 7.5c0 1.496-.913 3.255-2.688 3.652.838 1.699 1.438 3.768 1.181 5.697-.269 2.017-1.04 3.615-2.582 4.675C17.409 22.558 15.288 23 12.5 23H4.75a.75.75 0 010-1.5h2.322c-.58-.701-.998-1.578-1.223-2.471-.327-1.3-.297-2.786.265-4.131-.92.091-1.985-.02-3.126-.445a.75.75 0 11.524-1.406c1.964.733 3.428.266 4.045-.19.068-.06.137-.12.208-.18a.745.745 0 01.861-.076.746.746 0 01.32.368.752.752 0 01-.173.819c-.077.076-.16.15-.252.221-1.322 1.234-1.62 3.055-1.218 4.654.438 1.737 1.574 2.833 2.69 2.837H12.5c2.674 0 4.429-.433 5.56-1.212 1.094-.752 1.715-1.904 1.946-3.637.236-1.768-.445-3.845-1.407-5.529a.576.576 0 01-.012-.02 3.557 3.557 0 01-1.553-.94c-.556-.565-.89-1.243-1.012-1.73a.75.75 0 011.456-.364c.057.231.26.67.626 1.043.35.357.822.623 1.443.623 1.172 0 1.953-1.058 1.953-2.234 0-1.205-.357-2.127-.903-2.78-.547-.654-1.318-1.08-2.22-1.23z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Stack returns a string representing an "stack" Octicon.
func Stack(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-stack%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.122.392a1.75 1.75 0 011.756 0l5.003 2.902c.83.481.83 1.68 0 2.162L8.878 8.358a1.75 1.75 0 01-1.756 0L2.119 5.456a1.25 1.25 0 010-2.162L7.122.392zM8.125 1.69a.25.25 0 00-.25 0l-4.63 2.685 4.63 2.685a.25.25 0 00.25 0l4.63-2.685-4.63-2.685zM1.601 7.789a.75.75 0 011.025-.273l5.249 3.044a.25.25 0 00.25 0l5.249-3.044a.75.75 0 01.752 1.298l-5.248 3.044a1.75 1.75 0 01-1.756 0L1.874 8.814A.75.75 0 011.6 7.789zm0 3.5a.75.75 0 011.025-.273l5.249 3.044a.25.25 0 00.25 0l5.249-3.044a.75.75 0 01.752 1.298l-5.248 3.044a1.75 1.75 0 01-1.756 0l-5.248-3.044a.75.75 0 01-.273-1.025z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-stack%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.063 1.456a1.75 1.75 0 011.874 0l8.383 5.316a1.75 1.75 0 010 2.956l-8.383 5.316a1.75 1.75 0 01-1.874 0L2.68 9.728a1.75 1.75 0 010-2.956l8.383-5.316zm1.071 1.267a.25.25 0 00-.268 0L3.483 8.039a.25.25 0 000 .422l8.383 5.316a.25.25 0 00.268 0l8.383-5.316a.25.25 0 000-.422l-8.383-5.316z\"></path><path fill-rule=\"evenodd\" d=\"M1.867 12.324a.75.75 0 011.035-.232l8.964 5.685a.25.25 0 00.268 0l8.964-5.685a.75.75 0 01.804 1.267l-8.965 5.685a1.75 1.75 0 01-1.874 0l-8.965-5.685a.75.75 0 01-.231-1.035z\"></path><path fill-rule=\"evenodd\" d=\"M1.867 16.324a.75.75 0 011.035-.232l8.964 5.685a.25.25 0 00.268 0l8.964-5.685a.75.75 0 01.804 1.267l-8.965 5.685a1.75 1.75 0 01-1.874 0l-8.965-5.685a.75.75 0 01-.231-1.035z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Star returns a string representing an "star" Octicon.
func Star(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-star%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 .25a.75.75 0 01.673.418l1.882 3.815 4.21.612a.75.75 0 01.416 1.279l-3.046 2.97.719 4.192a.75.75 0 01-1.088.791L8 12.347l-3.766 1.98a.75.75 0 01-1.088-.79l.72-4.194L.818 6.374a.75.75 0 01.416-1.28l4.21-.611L7.327.668A.75.75 0 018 .25zm0 2.445L6.615 5.5a.75.75 0 01-.564.41l-3.097.45 2.24 2.184a.75.75 0 01.216.664l-.528 3.084 2.769-1.456a.75.75 0 01.698 0l2.77 1.456-.53-3.084a.75.75 0 01.216-.664l2.24-2.183-3.096-.45a.75.75 0 01-.564-.41L8 2.694v.001z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-star%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 .25a.75.75 0 01.673.418l3.058 6.197 6.839.994a.75.75 0 01.415 1.279l-4.948 4.823 1.168 6.811a.75.75 0 01-1.088.791L12 18.347l-6.117 3.216a.75.75 0 01-1.088-.79l1.168-6.812-4.948-4.823a.75.75 0 01.416-1.28l6.838-.993L11.328.668A.75.75 0 0112 .25zm0 2.445L9.44 7.882a.75.75 0 01-.565.41l-5.725.832 4.143 4.038a.75.75 0 01.215.664l-.978 5.702 5.121-2.692a.75.75 0 01.698 0l5.12 2.692-.977-5.702a.75.75 0 01.215-.664l4.143-4.038-5.725-.831a.75.75 0 01-.565-.41L12 2.694z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// StarFill returns a string representing an "star-fill" Octicon.
func StarFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-star-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 .25a.75.75 0 01.673.418l1.882 3.815 4.21.612a.75.75 0 01.416 1.279l-3.046 2.97.719 4.192a.75.75 0 01-1.088.791L8 12.347l-3.766 1.98a.75.75 0 01-1.088-.79l.72-4.194L.818 6.374a.75.75 0 01.416-1.28l4.21-.611L7.327.668A.75.75 0 018 .25z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-star-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.672.668a.75.75 0 00-1.345 0L8.27 6.865l-6.838.994a.75.75 0 00-.416 1.279l4.948 4.823-1.168 6.811a.75.75 0 001.088.791L12 18.347l6.117 3.216a.75.75 0 001.088-.79l-1.168-6.812 4.948-4.823a.75.75 0 00-.416-1.28l-6.838-.993L12.672.668z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Stop returns a string representing an "stop" Octicon.
func Stop(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-stop%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.47.22A.75.75 0 015 0h6a.75.75 0 01.53.22l4.25 4.25c.141.14.22.331.22.53v6a.75.75 0 01-.22.53l-4.25 4.25A.75.75 0 0111 16H5a.75.75 0 01-.53-.22L.22 11.53A.75.75 0 010 11V5a.75.75 0 01.22-.53L4.47.22zm.84 1.28L1.5 5.31v5.38l3.81 3.81h5.38l3.81-3.81V5.31L10.69 1.5H5.31zM8 4a.75.75 0 01.75.75v3.5a.75.75 0 01-1.5 0v-3.5A.75.75 0 018 4zm0 8a1 1 0 100-2 1 1 0 000 2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-stop%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12 7a.75.75 0 01.75.75v4.5a.75.75 0 01-1.5 0v-4.5A.75.75 0 0112 7zm0 10a1 1 0 100-2 1 1 0 000 2z\"></path><path fill-rule=\"evenodd\" d=\"M7.328 1.47a.75.75 0 01.53-.22h8.284a.75.75 0 01.53.22l5.858 5.858c.141.14.22.33.22.53v8.284a.75.75 0 01-.22.53l-5.858 5.858a.75.75 0 01-.53.22H7.858a.75.75 0 01-.53-.22L1.47 16.672a.75.75 0 01-.22-.53V7.858a.75.75 0 01.22-.53L7.328 1.47zm.84 1.28L2.75 8.169v7.662l5.419 5.419h7.662l5.419-5.418V8.168L15.832 2.75H8.168z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Stopwatch returns a string representing an "stopwatch" Octicon.
func Stopwatch(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-stopwatch%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.75.75A.75.75 0 016.5 0h3a.75.75 0 010 1.5h-.75v1l-.001.041a6.718 6.718 0 013.464 1.435l.007-.006.75-.75a.75.75 0 111.06 1.06l-.75.75-.006.007a6.75 6.75 0 11-10.548 0L2.72 5.03l-.75-.75a.75.75 0 011.06-1.06l.75.75.007.006A6.718 6.718 0 017.25 2.541a.756.756 0 010-.041v-1H6.5a.75.75 0 01-.75-.75zM8 14.5A5.25 5.25 0 108 4a5.25 5.25 0 000 10.5zm.389-6.7l1.33-1.33a.75.75 0 111.061 1.06L9.45 8.861A1.502 1.502 0 018 10.75a1.5 1.5 0 11.389-2.95z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-stopwatch%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.25 0a.75.75 0 000 1.5h1v1.278a9.955 9.955 0 00-5.635 2.276L4.28 3.72a.75.75 0 00-1.06 1.06l1.315 1.316A9.962 9.962 0 002 12.75c0 5.523 4.477 10 10 10s10-4.477 10-10a9.962 9.962 0 00-2.535-6.654L20.78 4.78a.75.75 0 00-1.06-1.06l-1.334 1.334a9.955 9.955 0 00-5.636-2.276V1.5h1a.75.75 0 000-1.5h-3.5zM12 21.25a8.5 8.5 0 100-17 8.5 8.5 0 000 17zm4.03-12.53a.75.75 0 010 1.06l-2.381 2.382a1.75 1.75 0 11-1.06-1.06l2.38-2.382a.75.75 0 011.061 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Strikethrough returns a string representing an "strikethrough" Octicon.
func Strikethrough(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-strikethrough%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.581 3.25c-2.036 0-2.778 1.082-2.778 1.786 0 .055.002.107.006.157a.75.75 0 01-1.496.114 3.56 3.56 0 01-.01-.271c0-1.832 1.75-3.286 4.278-3.286 1.418 0 2.721.58 3.514 1.093a.75.75 0 11-.814 1.26c-.64-.414-1.662-.853-2.7-.853zm3.474 5.25h3.195a.75.75 0 000-1.5H1.75a.75.75 0 000 1.5h6.018c.835.187 1.503.464 1.951.81.439.34.647.725.647 1.197 0 .428-.159.895-.594 1.267-.444.38-1.254.726-2.676.726-1.373 0-2.38-.493-2.86-.956a.75.75 0 00-1.042 1.079C3.992 13.393 5.39 14 7.096 14c1.652 0 2.852-.403 3.65-1.085a3.134 3.134 0 001.12-2.408 2.85 2.85 0 00-.811-2.007z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-strikethrough%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.36 5C9.37 5 8.105 6.613 8.105 7.848c0 .411.072.744.193 1.02a.75.75 0 01-1.373.603 3.993 3.993 0 01-.32-1.623c0-2.363 2.271-4.348 5.755-4.348 1.931 0 3.722.794 4.814 1.5a.75.75 0 11-.814 1.26c-.94-.607-2.448-1.26-4-1.26zm4.173 7.5h3.717a.75.75 0 000-1.5H3.75a.75.75 0 000 1.5h9.136c1.162.28 2.111.688 2.76 1.211.642.518.979 1.134.979 1.898a2.63 2.63 0 01-.954 2.036c-.703.601-1.934 1.105-3.999 1.105-2.018 0-3.529-.723-4.276-1.445a.75.75 0 10-1.042 1.08c1.066 1.028 2.968 1.865 5.318 1.865 2.295 0 3.916-.56 4.974-1.464a4.131 4.131 0 001.479-3.177c0-1.296-.608-2.316-1.538-3.066a5.77 5.77 0 00-.054-.043z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Sun returns a string representing an "sun" Octicon.
func Sun(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sun%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 10.5a2.5 2.5 0 100-5 2.5 2.5 0 000 5zM8 12a4 4 0 100-8 4 4 0 000 8zM8 0a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0V.75A.75.75 0 018 0zm0 13a.75.75 0 01.75.75v1.5a.75.75 0 01-1.5 0v-1.5A.75.75 0 018 13zM2.343 2.343a.75.75 0 011.061 0l1.06 1.061a.75.75 0 01-1.06 1.06l-1.06-1.06a.75.75 0 010-1.06zm9.193 9.193a.75.75 0 011.06 0l1.061 1.06a.75.75 0 01-1.06 1.061l-1.061-1.06a.75.75 0 010-1.061zM16 8a.75.75 0 01-.75.75h-1.5a.75.75 0 010-1.5h1.5A.75.75 0 0116 8zM3 8a.75.75 0 01-.75.75H.75a.75.75 0 010-1.5h1.5A.75.75 0 013 8zm10.657-5.657a.75.75 0 010 1.061l-1.061 1.06a.75.75 0 11-1.06-1.06l1.06-1.06a.75.75 0 011.06 0zm-9.193 9.193a.75.75 0 010 1.06l-1.06 1.061a.75.75 0 11-1.061-1.06l1.06-1.061a.75.75 0 011.061 0z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sun%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 17.5a5.5 5.5 0 100-11 5.5 5.5 0 000 11zm0 1.5a7 7 0 100-14 7 7 0 000 14zm12-7a.75.75 0 01-.75.75h-2.5a.75.75 0 010-1.5h2.5A.75.75 0 0124 12zM4 12a.75.75 0 01-.75.75H.75a.75.75 0 010-1.5h2.5A.75.75 0 014 12zm16.485-8.485a.75.75 0 010 1.06l-1.768 1.768a.75.75 0 01-1.06-1.06l1.767-1.768a.75.75 0 011.061 0zM6.343 17.657a.75.75 0 010 1.06l-1.768 1.768a.75.75 0 11-1.06-1.06l1.767-1.768a.75.75 0 011.061 0zM12 0a.75.75 0 01.75.75v2.5a.75.75 0 01-1.5 0V.75A.75.75 0 0112 0zm0 20a.75.75 0 01.75.75v2.5a.75.75 0 01-1.5 0v-2.5A.75.75 0 0112 20zM3.515 3.515a.75.75 0 011.06 0l1.768 1.768a.75.75 0 11-1.06 1.06L3.515 4.575a.75.75 0 010-1.06zm14.142 14.142a.75.75 0 011.06 0l1.768 1.768a.75.75 0 01-1.06 1.06l-1.768-1.767a.75.75 0 010-1.061z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Sync returns a string representing an "sync" Octicon.
func Sync(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-sync%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8 2.5a5.487 5.487 0 00-4.131 1.869l1.204 1.204A.25.25 0 014.896 6H1.25A.25.25 0 011 5.75V2.104a.25.25 0 01.427-.177l1.38 1.38A7.001 7.001 0 0114.95 7.16a.75.75 0 11-1.49.178A5.501 5.501 0 008 2.5zM1.705 8.005a.75.75 0 01.834.656 5.501 5.501 0 009.592 2.97l-1.204-1.204a.25.25 0 01.177-.427h3.646a.25.25 0 01.25.25v3.646a.25.25 0 01-.427.177l-1.38-1.38A7.001 7.001 0 011.05 8.84a.75.75 0 01.656-.834z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-sync%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M3.38 8A9.502 9.502 0 0112 2.5a9.502 9.502 0 019.215 7.182.75.75 0 101.456-.364C21.473 4.539 17.15 1 12 1a10.995 10.995 0 00-9.5 5.452V4.75a.75.75 0 00-1.5 0V8.5a1 1 0 001 1h3.75a.75.75 0 000-1.5H3.38zm-.595 6.318a.75.75 0 00-1.455.364C2.527 19.461 6.85 23 12 23c4.052 0 7.592-2.191 9.5-5.451v1.701a.75.75 0 001.5 0V15.5a1 1 0 00-1-1h-3.75a.75.75 0 000 1.5h2.37A9.502 9.502 0 0112 21.5c-4.446 0-8.181-3.055-9.215-7.182z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Tab returns a string representing an "tab" Octicon.
func Tab(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-tab%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M22 4.25a.75.75 0 00-1.5 0v15a.75.75 0 001.5 0v-15zm-9.72 14.28a.75.75 0 11-1.06-1.06l4.97-4.97H1.75a.75.75 0 010-1.5h14.44l-4.97-4.97a.75.75 0 011.06-1.06l6.25 6.25a.75.75 0 010 1.06l-6.25 6.25z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// TabExternal returns a string representing an "tab-external" Octicon.
func TabExternal(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-tab-external%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M3.25 4a.25.25 0 00-.25.25v9a.75.75 0 01-.75.75H.75a.75.75 0 010-1.5h.75V4.25c0-.966.784-1.75 1.75-1.75h9.5c.966 0 1.75.784 1.75 1.75v8.25h.75a.75.75 0 010 1.5h-1.5a.75.75 0 01-.75-.75v-9a.25.25 0 00-.25-.25h-9.5z\"></path><path d=\"M7.97 7.97l-2.75 2.75a.75.75 0 101.06 1.06l2.75-2.75 1.543 1.543a.25.25 0 00.427-.177V6.25a.25.25 0 00-.25-.25H6.604a.25.25 0 00-.177.427L7.97 7.97z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Table returns a string representing an "table" Octicon.
func Table(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-table%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v3.585a.746.746 0 010 .83v8.085A1.75 1.75 0 0114.25 16H6.309a.748.748 0 01-1.118 0H1.75A1.75 1.75 0 010 14.25V6.165a.746.746 0 010-.83V1.75zM1.5 6.5v7.75c0 .138.112.25.25.25H5v-8H1.5zM5 5H1.5V1.75a.25.25 0 01.25-.25H5V5zm1.5 1.5v8h7.75a.25.25 0 00.25-.25V6.5h-8zm8-1.5h-8V1.5h7.75a.25.25 0 01.25.25V5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-table%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2 3.75C2 2.784 2.784 2 3.75 2h16.5c.966 0 1.75.784 1.75 1.75v16.5A1.75 1.75 0 0120.25 22H3.75A1.75 1.75 0 012 20.25V3.75zM3.5 9v11.25c0 .138.112.25.25.25H7.5V9h-4zm4-1.5h-4V3.75a.25.25 0 01.25-.25H7.5v4zM9 9v11.5h11.25a.25.25 0 00.25-.25V9H9zm11.5-1.5H9v-4h11.25a.25.25 0 01.25.25V7.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Tag returns a string representing an "tag" Octicon.
func Tag(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-tag%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.5 7.775V2.75a.25.25 0 01.25-.25h5.025a.25.25 0 01.177.073l6.25 6.25a.25.25 0 010 .354l-5.025 5.025a.25.25 0 01-.354 0l-6.25-6.25a.25.25 0 01-.073-.177zm-1.5 0V2.75C1 1.784 1.784 1 2.75 1h5.025c.464 0 .91.184 1.238.513l6.25 6.25a1.75 1.75 0 010 2.474l-5.026 5.026a1.75 1.75 0 01-2.474 0l-6.25-6.25A1.75 1.75 0 011 7.775zM6 5a1 1 0 100 2 1 1 0 000-2z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-tag%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M7.75 6.5a1.25 1.25 0 100 2.5 1.25 1.25 0 000-2.5z\"></path><path fill-rule=\"evenodd\" d=\"M2.5 1A1.5 1.5 0 001 2.5v8.44c0 .397.158.779.44 1.06l10.25 10.25a1.5 1.5 0 002.12 0l8.44-8.44a1.5 1.5 0 000-2.12L12 1.44A1.5 1.5 0 0010.94 1H2.5zm0 1.5h8.44l10.25 10.25-8.44 8.44L2.5 10.94V2.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Tasklist returns a string representing an "tasklist" Octicon.
func Tasklist(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-tasklist%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.5 2.75a.25.25 0 01.25-.25h10.5a.25.25 0 01.25.25v10.5a.25.25 0 01-.25.25H2.75a.25.25 0 01-.25-.25V2.75zM2.75 1A1.75 1.75 0 001 2.75v10.5c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0015 13.25V2.75A1.75 1.75 0 0013.25 1H2.75zm9.03 5.28a.75.75 0 00-1.06-1.06L6.75 9.19 5.28 7.72a.75.75 0 00-1.06 1.06l2 2a.75.75 0 001.06 0l4.5-4.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-tasklist%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M17.28 9.28a.75.75 0 00-1.06-1.06l-5.97 5.97-2.47-2.47a.75.75 0 00-1.06 1.06l3 3a.75.75 0 001.06 0l6.5-6.5z\"></path><path fill-rule=\"evenodd\" d=\"M3.75 2A1.75 1.75 0 002 3.75v16.5c0 .966.784 1.75 1.75 1.75h16.5A1.75 1.75 0 0022 20.25V3.75A1.75 1.75 0 0020.25 2H3.75zM3.5 3.75a.25.25 0 01.25-.25h16.5a.25.25 0 01.25.25v16.5a.25.25 0 01-.25.25H3.75a.25.25 0 01-.25-.25V3.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Telescope returns a string representing an "telescope" Octicon.
func Telescope(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-telescope%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M14.184 1.143a1.75 1.75 0 00-2.502-.57L.912 7.916a1.75 1.75 0 00-.53 2.32l.447.775a1.75 1.75 0 002.275.702l11.745-5.656a1.75 1.75 0 00.757-2.451l-1.422-2.464zm-1.657.669a.25.25 0 01.358.081l1.422 2.464a.25.25 0 01-.108.35l-2.016.97-1.505-2.605 1.85-1.26zM9.436 3.92l1.391 2.41-5.42 2.61-.942-1.63 4.97-3.39zM3.222 8.157l-1.466 1a.25.25 0 00-.075.33l.447.775a.25.25 0 00.325.1l1.598-.769-.83-1.436zm6.253 2.306a.75.75 0 00-.944-.252l-1.809.87a.75.75 0 00-.293.253L4.38 14.326a.75.75 0 101.238.848l1.881-2.75v2.826a.75.75 0 001.5 0v-2.826l1.881 2.75a.75.75 0 001.238-.848l-2.644-3.863z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-telescope%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M.408 15.13a2 2 0 01.59-2.642L17.038 1.33a2 2 0 012.85.602l2.828 4.644a2 2 0 01-.851 2.847l-17.762 8.43a2 2 0 01-2.59-.807L.408 15.13zm5.263-4.066l7.842-5.455 2.857 4.76-8.712 4.135-1.987-3.44zm-1.235.86L1.854 13.72a.5.5 0 00-.147.66l1.105 1.915a.5.5 0 00.648.201l2.838-1.347-1.862-3.225zm13.295-2.2L14.747 4.75l3.148-2.19a.5.5 0 01.713.151l2.826 4.644a.5.5 0 01-.212.712l-3.49 1.656z\"></path><path d=\"M17.155 22.87a.75.75 0 00.226-1.036l-4-6.239a.75.75 0 00-.941-.278l-2.75 1.25a.75.75 0 00-.318.274l-3.25 4.989a.75.75 0 001.256.819l3.131-4.806.51-.232v5.64a.75.75 0 101.5 0v-6.22l3.6 5.613a.75.75 0 001.036.226z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// TelescopeFill returns a string representing an "telescope-fill" Octicon.
func TelescopeFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-telescope-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.531 10.21a.75.75 0 01.944.253l2.644 3.864a.75.75 0 11-1.238.847L9 12.424v2.826a.75.75 0 01-1.5 0v-2.826l-1.881 2.75a.75.75 0 01-1.238-.848l2.048-2.992a.75.75 0 01.293-.252l1.81-.871zM11.905.42a1.5 1.5 0 012.144.49l1.692 2.93a1.5 1.5 0 01-.649 2.102L2.895 11.815a1.5 1.5 0 01-1.95-.602l-.68-1.176a1.5 1.5 0 01.455-1.99L11.905.422zM3.279 8.119l.835 1.445 1.355-.653-.947-1.64-1.243.848zm7.728-1.874L9.6 3.808l1.243-.848 1.52 2.631-1.356.653z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-telescope-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M17.155 22.87a.75.75 0 00.226-1.036l-4-6.239a.75.75 0 00-.941-.277l-2.75 1.25a.75.75 0 00-.318.273l-3.25 4.989a.75.75 0 001.256.819l3.131-4.806.51-.232v5.64a.75.75 0 101.5 0v-6.22l3.6 5.613a.75.75 0 001.036.226z\"></path><path fill-rule=\"evenodd\" d=\"M.408 15.13a2 2 0 01.59-2.642L17.038 1.33a2 2 0 012.85.602l2.828 4.644a2 2 0 01-.851 2.847l-17.762 8.43a2 2 0 01-2.59-.807L.408 15.13zm5.263-4.066l1.987 3.44-1.36.645-1.862-3.225 1.235-.86zm7.842-5.455l2.857 4.76 1.361-.646-2.984-4.973-1.234.859z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Terminal returns a string representing an "terminal" Octicon.
func Terminal(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-terminal%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 2.75C0 1.784.784 1 1.75 1h12.5c.966 0 1.75.784 1.75 1.75v10.5A1.75 1.75 0 0114.25 15H1.75A1.75 1.75 0 010 13.25V2.75zm1.75-.25a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25V2.75a.25.25 0 00-.25-.25H1.75zM7.25 8a.75.75 0 01-.22.53l-2.25 2.25a.75.75 0 11-1.06-1.06L5.44 8 3.72 6.28a.75.75 0 111.06-1.06l2.25 2.25c.141.14.22.331.22.53zm1.5 1.5a.75.75 0 000 1.5h3a.75.75 0 000-1.5h-3z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-terminal%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M9.25 12a.75.75 0 01-.22.53l-2.75 2.75a.75.75 0 01-1.06-1.06L7.44 12 5.22 9.78a.75.75 0 111.06-1.06l2.75 2.75c.141.14.22.331.22.53zm2 2a.75.75 0 000 1.5h5a.75.75 0 000-1.5h-5z\"></path><path fill-rule=\"evenodd\" d=\"M0 4.75C0 3.784.784 3 1.75 3h20.5c.966 0 1.75.784 1.75 1.75v14.5A1.75 1.75 0 0122.25 21H1.75A1.75 1.75 0 010 19.25V4.75zm1.75-.25a.25.25 0 00-.25.25v14.5c0 .138.112.25.25.25h20.5a.25.25 0 00.25-.25V4.75a.25.25 0 00-.25-.25H1.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// ThreeBars returns a string representing an "three-bars" Octicon.
func ThreeBars(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-three-bars%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1 2.75A.75.75 0 011.75 2h12.5a.75.75 0 110 1.5H1.75A.75.75 0 011 2.75zm0 5A.75.75 0 011.75 7h12.5a.75.75 0 110 1.5H1.75A.75.75 0 011 7.75zM1.75 12a.75.75 0 100 1.5h12.5a.75.75 0 100-1.5H1.75z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Thumbsdown returns a string representing an "thumbsdown" Octicon.
func Thumbsdown(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-thumbsdown%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.083 15.986c1.34.153 2.334-.982 2.334-2.183v-.5c0-1.329.646-2.123 1.317-2.614.329-.24.66-.403.919-.508a1.75 1.75 0 001.514.872h1a1.75 1.75 0 001.75-1.75v-7.5a1.75 1.75 0 00-1.75-1.75h-1a1.75 1.75 0 00-1.662 1.2c-.525-.074-1.068-.228-1.726-.415L9.305.705C8.151.385 6.765.053 4.917.053c-1.706 0-2.97.152-3.722 1.139-.353.463-.537 1.042-.669 1.672C.41 3.424.32 4.108.214 4.897l-.04.306c-.25 1.869-.266 3.318.188 4.316.244.537.622.943 1.136 1.2.495.248 1.066.334 1.669.334h1.422l-.015.112c-.07.518-.157 1.17-.157 1.638 0 .921.151 1.718.655 2.299.512.589 1.248.797 2.011.884zm4.334-13.232c-.706-.089-1.39-.284-2.072-.479a63.914 63.914 0 00-.441-.125c-1.096-.304-2.335-.597-3.987-.597-1.794 0-2.28.222-2.529.548-.147.193-.275.505-.393 1.07-.105.502-.188 1.124-.295 1.93l-.04.3c-.25 1.882-.19 2.933.067 3.497a.921.921 0 00.443.48c.208.104.52.175.997.175h1.75c.685 0 1.295.577 1.205 1.335-.022.192-.049.39-.075.586-.066.488-.13.97-.13 1.329 0 .808.144 1.15.288 1.316.137.157.401.303 1.048.377.307.035.664-.237.664-.693v-.5c0-1.922.978-3.127 1.932-3.825a5.862 5.862 0 011.568-.809V2.754zm1.75 6.798a.25.25 0 01-.25-.25v-7.5a.25.25 0 01.25-.25h1a.25.25 0 01.25.25v7.5a.25.25 0 01-.25.25h-1z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-thumbsdown%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.596 21.957c-1.301.092-2.303-.986-2.303-2.206v-1.053c0-2.666-1.813-3.785-2.774-4.2a1.864 1.864 0 00-.523-.13A1.75 1.75 0 015.25 16h-1.5A1.75 1.75 0 012 14.25V3.75C2 2.784 2.784 2 3.75 2h1.5a1.75 1.75 0 011.742 1.58c.838-.06 1.667-.296 2.69-.586l.602-.17C11.748 2.419 13.497 2 15.828 2c2.188 0 3.693.204 4.583 1.372.422.554.65 1.255.816 2.05.148.708.262 1.57.396 2.58l.051.39c.319 2.386.328 4.18-.223 5.394-.293.644-.743 1.125-1.355 1.431-.59.296-1.284.404-2.036.404h-2.05l.056.429c.025.18.05.372.076.572.06.483.117 1.006.117 1.438 0 1.245-.222 2.253-.92 2.942-.684.674-1.668.879-2.743.955zM7 5.082c1.059-.064 2.079-.355 3.118-.651.188-.054.377-.108.568-.16 1.406-.392 3.006-.771 5.142-.771 2.277 0 3.004.274 3.39.781.216.283.388.718.54 1.448.136.65.242 1.45.379 2.477l.05.385c.32 2.398.253 3.794-.102 4.574-.16.352-.375.569-.66.711-.305.153-.74.245-1.365.245h-2.37c-.681 0-1.293.57-1.211 1.328.026.244.065.537.105.834l.07.527c.06.482.105.922.105 1.25 0 1.125-.213 1.617-.473 1.873-.275.27-.774.456-1.795.528-.351.024-.698-.274-.698-.71v-1.053c0-3.55-2.488-5.063-3.68-5.577A3.485 3.485 0 007 12.861V5.08zM3.75 3.5a.25.25 0 00-.25.25v10.5c0 .138.112.25.25.25h1.5a.25.25 0 00.25-.25V3.75a.25.25 0 00-.25-.25h-1.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Thumbsup returns a string representing an "thumbsup" Octicon.
func Thumbsup(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-thumbsup%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.834.066C7.494-.087 6.5 1.048 6.5 2.25v.5c0 1.329-.647 2.124-1.318 2.614-.328.24-.66.403-.918.508A1.75 1.75 0 002.75 5h-1A1.75 1.75 0 000 6.75v7.5C0 15.216.784 16 1.75 16h1a1.75 1.75 0 001.662-1.201c.525.075 1.067.229 1.725.415.152.043.31.088.475.133 1.154.32 2.54.653 4.388.653 1.706 0 2.97-.153 3.722-1.14.353-.463.537-1.042.668-1.672.118-.56.208-1.243.313-2.033l.04-.306c.25-1.869.265-3.318-.188-4.316a2.418 2.418 0 00-1.137-1.2C13.924 5.085 13.353 5 12.75 5h-1.422l.015-.113c.07-.518.157-1.17.157-1.637 0-.922-.151-1.719-.656-2.3-.51-.589-1.247-.797-2.01-.884zM4.5 13.3c.705.088 1.39.284 2.072.478l.441.125c1.096.305 2.334.598 3.987.598 1.794 0 2.28-.223 2.528-.549.147-.193.276-.505.394-1.07.105-.502.188-1.124.295-1.93l.04-.3c.25-1.882.189-2.933-.068-3.497a.922.922 0 00-.442-.48c-.208-.104-.52-.174-.997-.174H11c-.686 0-1.295-.577-1.206-1.336.023-.192.05-.39.076-.586.065-.488.13-.97.13-1.328 0-.809-.144-1.15-.288-1.316-.137-.158-.402-.304-1.048-.378C8.357 1.521 8 1.793 8 2.25v.5c0 1.922-.978 3.128-1.933 3.825a5.861 5.861 0 01-1.567.81V13.3zM2.75 6.5a.25.25 0 01.25.25v7.5a.25.25 0 01-.25.25h-1a.25.25 0 01-.25-.25v-7.5a.25.25 0 01.25-.25h1z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-thumbsup%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12.596 2.043c-1.301-.092-2.303.986-2.303 2.206v1.053c0 2.666-1.813 3.785-2.774 4.2a1.866 1.866 0 01-.523.131A1.75 1.75 0 005.25 8h-1.5A1.75 1.75 0 002 9.75v10.5c0 .967.784 1.75 1.75 1.75h1.5a1.75 1.75 0 001.742-1.58c.838.06 1.667.296 2.69.586l.602.17c1.464.406 3.213.824 5.544.824 2.188 0 3.693-.204 4.583-1.372.422-.554.65-1.255.816-2.05.148-.708.262-1.57.396-2.58l.051-.39c.319-2.386.328-4.18-.223-5.394-.293-.644-.743-1.125-1.355-1.431-.59-.296-1.284-.404-2.036-.404h-2.05l.056-.429c.025-.18.05-.372.076-.572.06-.483.117-1.006.117-1.438 0-1.245-.222-2.253-.92-2.941-.684-.675-1.668-.88-2.743-.956zM7 18.918c1.059.064 2.079.355 3.118.652l.568.16c1.406.39 3.006.77 5.142.77 2.277 0 3.004-.274 3.39-.781.216-.283.388-.718.54-1.448.136-.65.242-1.45.379-2.477l.05-.384c.32-2.4.253-3.795-.102-4.575-.16-.352-.375-.568-.66-.711-.305-.153-.74-.245-1.365-.245h-2.37c-.681 0-1.293-.57-1.211-1.328.026-.243.065-.537.105-.834l.07-.527c.06-.482.105-.921.105-1.25 0-1.125-.213-1.617-.473-1.873-.275-.27-.774-.455-1.795-.528-.351-.024-.698.274-.698.71v1.053c0 3.55-2.488 5.063-3.68 5.577-.372.16-.754.232-1.113.26v7.78zM3.75 20.5a.25.25 0 01-.25-.25V9.75a.25.25 0 01.25-.25h1.5a.25.25 0 01.25.25v10.5a.25.25 0 01-.25.25h-1.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Tools returns a string representing an "tools" Octicon.
func Tools(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-tools%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.433 2.304A4.494 4.494 0 003.5 6c0 1.598.832 3.002 2.09 3.802.518.328.929.923.902 1.64v.008l-.164 3.337a.75.75 0 11-1.498-.073l.163-3.33c.002-.085-.05-.216-.207-.316A5.996 5.996 0 012 6a5.994 5.994 0 012.567-4.92 1.482 1.482 0 011.673-.04c.462.296.76.827.76 1.423v2.82c0 .082.041.16.11.206l.75.51a.25.25 0 00.28 0l.75-.51A.25.25 0 009 5.282V2.463c0-.596.298-1.127.76-1.423a1.482 1.482 0 011.673.04A5.994 5.994 0 0114 6a5.996 5.996 0 01-2.786 5.068c-.157.1-.209.23-.207.315l.163 3.33a.75.75 0 11-1.498.074l-.164-3.345c-.027-.717.384-1.312.902-1.64A4.496 4.496 0 0012.5 6a4.494 4.494 0 00-1.933-3.696c-.024.017-.067.067-.067.16v2.818a1.75 1.75 0 01-.767 1.448l-.75.51a1.75 1.75 0 01-1.966 0l-.75-.51A1.75 1.75 0 015.5 5.282V2.463c0-.092-.043-.142-.067-.159zm.01-.005z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-tools%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.875 2.292a.125.125 0 00-.032.018A7.24 7.24 0 004.75 8.25a7.247 7.247 0 003.654 6.297c.57.327.982.955.941 1.682v.002l-.317 6.058a.75.75 0 11-1.498-.078l.317-6.062v-.004c.006-.09-.047-.215-.188-.296A8.747 8.747 0 013.25 8.25a8.74 8.74 0 013.732-7.169 1.547 1.547 0 011.709-.064c.484.292.809.835.809 1.46v4.714a.25.25 0 00.119.213l2.25 1.385c.08.05.182.05.262 0l2.25-1.385a.25.25 0 00.119-.213V2.478c0-.626.325-1.169.81-1.461a1.547 1.547 0 011.708.064 8.74 8.74 0 013.732 7.17 8.747 8.747 0 01-4.41 7.598c-.14.081-.193.206-.188.296v.004l.318 6.062a.75.75 0 11-1.498.078l-.317-6.058v-.002c-.041-.727.37-1.355.94-1.682A7.247 7.247 0 0019.25 8.25a7.24 7.24 0 00-3.093-5.94.125.125 0 00-.032-.018l-.01-.001c-.003 0-.014 0-.031.01-.036.022-.084.079-.084.177V7.19a1.75 1.75 0 01-.833 1.49l-2.25 1.385a1.75 1.75 0 01-1.834 0l-2.25-1.384A1.75 1.75 0 018 7.192V2.477c0-.098-.048-.155-.084-.176a.062.062 0 00-.031-.011l-.01.001z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Trash returns a string representing an "trash" Octicon.
func Trash(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-trash%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.5 1.75a.25.25 0 01.25-.25h2.5a.25.25 0 01.25.25V3h-3V1.75zm4.5 0V3h2.25a.75.75 0 010 1.5H2.75a.75.75 0 010-1.5H5V1.75C5 .784 5.784 0 6.75 0h2.5C10.216 0 11 .784 11 1.75zM4.496 6.675a.75.75 0 10-1.492.15l.66 6.6A1.75 1.75 0 005.405 15h5.19c.9 0 1.652-.681 1.741-1.576l.66-6.6a.75.75 0 00-1.492-.149l-.66 6.6a.25.25 0 01-.249.225h-5.19a.25.25 0 01-.249-.225l-.66-6.6z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-trash%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M16 1.75V3h5.25a.75.75 0 010 1.5H2.75a.75.75 0 010-1.5H8V1.75C8 .784 8.784 0 9.75 0h4.5C15.216 0 16 .784 16 1.75zm-6.5 0a.25.25 0 01.25-.25h4.5a.25.25 0 01.25.25V3h-5V1.75z\"></path><path d=\"M4.997 6.178a.75.75 0 10-1.493.144L4.916 20.92a1.75 1.75 0 001.742 1.58h10.684a1.75 1.75 0 001.742-1.581l1.413-14.597a.75.75 0 00-1.494-.144l-1.412 14.596a.25.25 0 01-.249.226H6.658a.25.25 0 01-.249-.226L4.997 6.178z\"></path><path d=\"M9.206 7.501a.75.75 0 01.793.705l.5 8.5A.75.75 0 119 16.794l-.5-8.5a.75.75 0 01.705-.793zm6.293.793A.75.75 0 1014 8.206l-.5 8.5a.75.75 0 001.498.088l.5-8.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// TriangleDown returns a string representing an "triangle-down" Octicon.
func TriangleDown(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-down%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M4.427 7.427l3.396 3.396a.25.25 0 00.354 0l3.396-3.396A.25.25 0 0011.396 7H4.604a.25.25 0 00-.177.427z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-down%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M11.646 15.146L5.854 9.354a.5.5 0 01.353-.854h11.586a.5.5 0 01.353.854l-5.793 5.792a.5.5 0 01-.707 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// TriangleLeft returns a string representing an "triangle-left" Octicon.
func TriangleLeft(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-left%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M9.573 4.427L6.177 7.823a.25.25 0 000 .354l3.396 3.396a.25.25 0 00.427-.177V4.604a.25.25 0 00-.427-.177z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-left%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M8.854 11.646l5.792-5.792a.5.5 0 01.854.353v11.586a.5.5 0 01-.854.353l-5.792-5.792a.5.5 0 010-.708z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// TriangleRight returns a string representing an "triangle-right" Octicon.
func TriangleRight(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-right%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M6.427 4.427l3.396 3.396a.25.25 0 010 .354l-3.396 3.396A.25.25 0 016 11.396V4.604a.25.25 0 01.427-.177z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-right%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M15.146 12.354l-5.792 5.792a.5.5 0 01-.854-.353V6.207a.5.5 0 01.854-.353l5.792 5.792a.5.5 0 010 .708z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// TriangleUp returns a string representing an "triangle-up" Octicon.
func TriangleUp(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-up%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M4.427 9.573l3.396-3.396a.25.25 0 01.354 0l3.396 3.396a.25.25 0 01-.177.427H4.604a.25.25 0 01-.177-.427z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-triangle-up%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M12.354 8.854l5.792 5.792a.5.5 0 01-.353.854H6.207a.5.5 0 01-.353-.854l5.792-5.792a.5.5 0 01.708 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Trophy returns a string representing an "trophy" Octicon.
func Trophy(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-trophy%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.217 6.962A3.75 3.75 0 010 3.25v-.5C0 1.784.784 1 1.75 1h1.356c.228-.585.796-1 1.462-1h6.864a1.57 1.57 0 011.462 1h1.356c.966 0 1.75.784 1.75 1.75v.5a3.75 3.75 0 01-3.217 3.712 5.014 5.014 0 01-2.771 3.117l.144 1.446c.005.05.03.12.114.204.086.087.217.17.373.227.283.103.618.274.89.568.285.31.467.723.467 1.226v.75h1.25a.75.75 0 110 1.5H2.75a.75.75 0 010-1.5H4v-.75c0-.503.182-.916.468-1.226.27-.294.606-.465.889-.568a1.03 1.03 0 00.373-.227c.084-.085.109-.153.114-.204l.144-1.446a5.014 5.014 0 01-2.77-3.117zM3 2.5H1.75a.25.25 0 00-.25.25v.5c0 .98.626 1.813 1.5 2.122V2.5zm4.457 7.97l-.12 1.204c-.093.925-.858 1.47-1.467 1.691a.764.764 0 00-.3.176c-.037.04-.07.093-.07.21v.75h5v-.75c0-.117-.033-.17-.07-.21a.763.763 0 00-.3-.176c-.609-.221-1.374-.766-1.466-1.69l-.12-1.204a5.052 5.052 0 01-1.087 0zM13 5.373V2.5h1.25a.25.25 0 01.25.25v.5A2.25 2.25 0 0113 5.372zM4.5 1.568c0-.037.03-.068.068-.068h6.864c.037 0 .068.03.068.068V5.5a3.5 3.5 0 11-7 0V1.568z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-trophy%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.09 10.121A5.252 5.252 0 011 5V3.75C1 2.784 1.784 2 2.75 2h2.364c.236-.586.81-1 1.48-1h10.812c.67 0 1.244.414 1.48 1h2.489c.966 0 1.75.784 1.75 1.75V5a5.252 5.252 0 01-4.219 5.149 7.01 7.01 0 01-4.644 5.478l.231 3.003a.326.326 0 00.034.031c.079.065.303.203.836.282.838.124 1.637.81 1.637 1.807v.75h2.25a.75.75 0 010 1.5H4.75a.75.75 0 010-1.5H7v-.75c0-.996.8-1.683 1.637-1.807.533-.08.757-.217.836-.282a.334.334 0 00.034-.031l.231-3.003A7.01 7.01 0 015.09 10.12zM5 3.5H2.75a.25.25 0 00-.25.25V5A3.752 3.752 0 005 8.537V3.5zm6.217 12.457l-.215 2.793-.001.021-.003.043a1.203 1.203 0 01-.022.147c-.05.237-.194.567-.553.86-.348.286-.853.5-1.566.605a.482.482 0 00-.274.136.265.265 0 00-.083.188v.75h7v-.75a.265.265 0 00-.083-.188.483.483 0 00-.274-.136c-.713-.105-1.218-.32-1.567-.604-.358-.294-.502-.624-.552-.86a1.203 1.203 0 01-.025-.19l-.001-.022-.215-2.793a7.076 7.076 0 01-1.566 0zM19 8.578V3.5h2.375a.25.25 0 01.25.25V5c0 1.68-1.104 3.1-2.625 3.578zM6.5 2.594c0-.052.042-.094.094-.094h10.812c.052 0 .094.042.094.094V9a5.5 5.5 0 11-11 0V2.594z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Typography returns a string representing an "typography" Octicon.
func Typography(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-typography%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.21 8.5L4.574 3.594 2.857 8.5H6.21zm.5 1.5l.829 2.487a.75.75 0 001.423-.474L5.735 2.332a1.216 1.216 0 00-2.302-.018l-3.39 9.688a.75.75 0 001.415.496L2.332 10H6.71zm3.13-4.358C10.53 4.374 11.87 4 13 4c1.5 0 3 .939 3 2.601v5.649a.75.75 0 01-1.448.275C13.995 12.82 13.3 13 12.5 13c-.77 0-1.514-.231-2.078-.709-.577-.488-.922-1.199-.922-2.041 0-.694.265-1.411.887-1.944C11 7.78 11.88 7.5 13 7.5h1.5v-.899c0-.54-.5-1.101-1.5-1.101-.869 0-1.528.282-1.84.858a.75.75 0 11-1.32-.716zM14.5 9H13c-.881 0-1.375.22-1.637.444-.253.217-.363.5-.363.806 0 .408.155.697.39.896.249.21.63.354 1.11.354.732 0 1.26-.209 1.588-.449.35-.257.412-.495.412-.551V9z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-typography%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.414 15l1.63 4.505a.75.75 0 001.411-.51l-5.08-14.03a1.463 1.463 0 00-2.75 0l-5.08 14.03a.75.75 0 101.41.51L3.586 15h6.828zm-.544-1.5L7 5.572 4.13 13.5h5.74zm5.076-3.598c.913-1.683 2.703-2.205 4.284-2.205 1.047 0 2.084.312 2.878.885.801.577 1.392 1.455 1.392 2.548v8.12a.75.75 0 01-1.5 0v-.06a3.123 3.123 0 01-.044.025c-.893.52-2.096.785-3.451.785-1.051 0-2.048-.315-2.795-.948-.76-.643-1.217-1.578-1.217-2.702 0-.919.349-1.861 1.168-2.563.81-.694 2-1.087 3.569-1.087H22v-1.57c0-.503-.263-.967-.769-1.332-.513-.37-1.235-.6-2.001-.6-1.319 0-2.429.43-2.966 1.42a.75.75 0 01-1.318-.716zM22 14.2h-2.77c-1.331 0-2.134.333-2.593.726a1.82 1.82 0 00-.644 1.424c0 .689.267 1.203.686 1.557.43.365 1.065.593 1.826.593 1.183 0 2.102-.235 2.697-.581.582-.34.798-.74.798-1.134V14.2z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Unfold returns a string representing an "unfold" Octicon.
func Unfold(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-unfold%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M8.177.677l2.896 2.896a.25.25 0 01-.177.427H8.75v1.25a.75.75 0 01-1.5 0V4H5.104a.25.25 0 01-.177-.427L7.823.677a.25.25 0 01.354 0zM7.25 10.75a.75.75 0 011.5 0V12h2.146a.25.25 0 01.177.427l-2.896 2.896a.25.25 0 01-.354 0l-2.896-2.896A.25.25 0 015.104 12H7.25v-1.25zm-5-2a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5zM6 8a.75.75 0 01-.75.75h-.5a.75.75 0 010-1.5h.5A.75.75 0 016 8zm2.25.75a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5zM12 8a.75.75 0 01-.75.75h-.5a.75.75 0 010-1.5h.5A.75.75 0 0112 8zm2.25.75a.75.75 0 000-1.5h-.5a.75.75 0 000 1.5h.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-unfold%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M12 23a.75.75 0 01-.53-.22l-3.25-3.25a.75.75 0 111.06-1.06L12 21.19l2.72-2.72a.75.75 0 111.06 1.06l-3.25 3.25A.75.75 0 0112 23z\"></path><path fill-rule=\"evenodd\" d=\"M12 22.25a.75.75 0 01-.75-.75v-5.75a.75.75 0 011.5 0v5.75a.75.75 0 01-.75.75zM10.75 12a.75.75 0 01.75-.75h1a.75.75 0 110 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm-8 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zm12 0a.75.75 0 01.75-.75h1a.75.75 0 010 1.5h-1a.75.75 0 01-.75-.75zM11.47 1.22a.75.75 0 011.06 0l3.25 3.25a.75.75 0 01-1.06 1.06L12 2.81 9.28 5.53a.75.75 0 01-1.06-1.06l3.25-3.25z\"></path><path fill-rule=\"evenodd\" d=\"M12 1.5a.75.75 0 01.75.75v6a.75.75 0 01-1.5 0v-6A.75.75 0 0112 1.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Unlock returns a string representing an "unlock" Octicon.
func Unlock(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-unlock%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.5 4a2.5 2.5 0 014.607-1.346.75.75 0 101.264-.808A4 4 0 004 4v2h-.501A1.5 1.5 0 002 7.5v6A1.5 1.5 0 003.5 15h9a1.5 1.5 0 001.5-1.5v-6A1.5 1.5 0 0012.5 6h-7V4zm-.75 3.5H3.5v6h9v-6H4.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-unlock%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.5 7.25C7.5 4.58 9.422 2.5 12 2.5c2.079 0 3.71 1.34 4.282 3.242a.75.75 0 101.436-.432C16.971 2.825 14.792 1 12 1 8.503 1 6 3.845 6 7.25V9h-.5A2.5 2.5 0 003 11.5v8A2.5 2.5 0 005.5 22h13a2.5 2.5 0 002.5-2.5v-8A2.5 2.5 0 0018.5 9h-11V7.25zm-3 4.25a1 1 0 011-1h13a1 1 0 011 1v8a1 1 0 01-1 1h-13a1 1 0 01-1-1v-8z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Unmute returns a string representing an "unmute" Octicon.
func Unmute(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-unmute%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.563 2.069A.75.75 0 018 2.75v10.5a.75.75 0 01-1.238.57L3.472 11H1.75A1.75 1.75 0 010 9.25v-2.5C0 5.784.784 5 1.75 5h1.723l3.289-2.82a.75.75 0 01.801-.111zM6.5 4.38L4.238 6.319a.75.75 0 01-.488.181h-2a.25.25 0 00-.25.25v2.5c0 .138.112.25.25.25h2a.75.75 0 01.488.18L6.5 11.62V4.38zm6.096-2.038a.75.75 0 011.06 0 8 8 0 010 11.314.75.75 0 01-1.06-1.06 6.5 6.5 0 000-9.193.75.75 0 010-1.06v-.001zm-1.06 2.121a.75.75 0 10-1.061 1.061 3.5 3.5 0 010 4.95.75.75 0 101.06 1.06 5 5 0 000-7.07l.001-.001z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-unmute%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M11.553 3.064A.75.75 0 0112 3.75v16.5a.75.75 0 01-1.255.555L5.46 16H2.75A1.75 1.75 0 011 14.25v-4.5C1 8.784 1.784 8 2.75 8h2.71l5.285-4.805a.75.75 0 01.808-.13zM10.5 5.445l-4.245 3.86a.75.75 0 01-.505.195h-3a.25.25 0 00-.25.25v4.5c0 .138.112.25.25.25h3a.75.75 0 01.505.195l4.245 3.86V5.445z\"></path><path d=\"M18.718 4.222a.75.75 0 011.06 0c4.296 4.296 4.296 11.26 0 15.556a.75.75 0 01-1.06-1.06 9.5 9.5 0 000-13.436.75.75 0 010-1.06z\"></path><path d=\"M16.243 7.757a.75.75 0 10-1.061 1.061 4.5 4.5 0 010 6.364.75.75 0 001.06 1.06 6 6 0 000-8.485z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Unverified returns a string representing an "unverified" Octicon.
func Unverified(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-unverified%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M6.415.52a2.678 2.678 0 013.17 0l.928.68c.153.113.33.186.518.215l1.138.175a2.678 2.678 0 012.241 2.24l.175 1.138c.029.187.102.365.215.518l.68.928a2.678 2.678 0 010 3.17l-.68.928a1.179 1.179 0 00-.215.518l-.175 1.138a2.678 2.678 0 01-2.241 2.241l-1.138.175a1.179 1.179 0 00-.518.215l-.928.68a2.678 2.678 0 01-3.17 0l-.928-.68a1.179 1.179 0 00-.518-.215L3.83 14.41a2.678 2.678 0 01-2.24-2.24l-.175-1.138a1.179 1.179 0 00-.215-.518l-.68-.928a2.678 2.678 0 010-3.17l.68-.928a1.17 1.17 0 00.215-.518l.175-1.14a2.678 2.678 0 012.24-2.24l1.138-.175c.187-.029.365-.102.518-.215l.928-.68zm2.282 1.209a1.178 1.178 0 00-1.394 0l-.928.68a2.678 2.678 0 01-1.18.489l-1.136.174a1.178 1.178 0 00-.987.987l-.174 1.137a2.678 2.678 0 01-.489 1.18l-.68.927c-.305.415-.305.98 0 1.394l.68.928c.256.348.423.752.489 1.18l.174 1.136c.078.51.478.909.987.987l1.137.174c.427.066.831.233 1.18.489l.927.68c.415.305.98.305 1.394 0l.928-.68a2.678 2.678 0 011.18-.489l1.136-.174c.51-.078.909-.478.987-.987l.174-1.137c.066-.427.233-.831.489-1.18l.68-.927c.305-.415.305-.98 0-1.394l-.68-.928a2.678 2.678 0 01-.489-1.18l-.174-1.136a1.178 1.178 0 00-.987-.987l-1.137-.174a2.678 2.678 0 01-1.18-.489l-.927-.68zM9 11a1 1 0 11-2 0 1 1 0 012 0zM6.92 6.085c.081-.16.19-.299.34-.398.145-.097.371-.187.74-.187.28 0 .553.087.738.225A.613.613 0 019 6.25c0 .177-.04.264-.077.318a.956.956 0 01-.277.245c-.076.051-.158.1-.258.161l-.007.004c-.093.056-.204.122-.313.195a2.416 2.416 0 00-.692.661.75.75 0 001.248.832.956.956 0 01.276-.245 6.3 6.3 0 01.26-.16l.006-.004c.093-.057.204-.123.313-.195.222-.149.487-.355.692-.662.214-.32.329-.702.329-1.15 0-.76-.36-1.348-.862-1.725A2.76 2.76 0 008 4c-.631 0-1.154.16-1.572.438-.413.276-.68.638-.849.977a.75.75 0 001.342.67z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-unverified%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M13 16.5a1 1 0 11-2 0 1 1 0 012 0zm-2.517-7.665c.112-.223.268-.424.488-.57C11.186 8.12 11.506 8 12 8c.384 0 .766.118 1.034.319a.953.953 0 01.403.806c0 .48-.218.81-.62 1.186a9.293 9.293 0 01-.409.354 19.8 19.8 0 00-.294.249c-.246.213-.524.474-.738.795l-.126.19V13.5a.75.75 0 001.5 0v-1.12c.09-.1.203-.208.347-.333.063-.055.14-.119.222-.187.166-.14.358-.3.52-.452.536-.5 1.098-1.2 1.098-2.283a2.45 2.45 0 00-1.003-2.006C13.37 6.695 12.658 6.5 12 6.5c-.756 0-1.373.191-1.861.517a2.944 2.944 0 00-.997 1.148.75.75 0 001.341.67z\"></path><path fill-rule=\"evenodd\" d=\"M9.864 1.2a3.61 3.61 0 014.272 0l1.375 1.01c.274.2.593.333.929.384l1.686.259a3.61 3.61 0 013.021 3.02l.259 1.687c.051.336.183.655.384.929l1.01 1.375a3.61 3.61 0 010 4.272l-1.01 1.375a2.11 2.11 0 00-.384.929l-.259 1.686a3.61 3.61 0 01-3.02 3.021l-1.687.259a2.11 2.11 0 00-.929.384l-1.375 1.01a3.61 3.61 0 01-4.272 0l-1.375-1.01a2.11 2.11 0 00-.929-.384l-1.686-.259a3.61 3.61 0 01-3.021-3.02l-.259-1.687a2.11 2.11 0 00-.384-.929L1.2 14.136a3.61 3.61 0 010-4.272l1.01-1.375a2.11 2.11 0 00.384-.929l.259-1.686a3.61 3.61 0 013.02-3.021l1.687-.259a2.11 2.11 0 00.929-.384L9.864 1.2zm3.384 1.209a2.11 2.11 0 00-2.496 0l-1.376 1.01a3.61 3.61 0 01-1.589.658l-1.686.258a2.11 2.11 0 00-1.766 1.766l-.258 1.686a3.61 3.61 0 01-.658 1.59l-1.01 1.375a2.11 2.11 0 000 2.496l1.01 1.376a3.61 3.61 0 01.658 1.589l.258 1.686a2.11 2.11 0 001.766 1.765l1.686.26a3.61 3.61 0 011.59.657l1.375 1.01a2.11 2.11 0 002.496 0l1.376-1.01a3.61 3.61 0 011.589-.658l1.686-.258a2.11 2.11 0 001.765-1.766l.26-1.686a3.61 3.61 0 01.657-1.59l1.01-1.375a2.11 2.11 0 000-2.496l-1.01-1.376a3.61 3.61 0 01-.658-1.589l-.258-1.686a2.11 2.11 0 00-1.766-1.766l-1.686-.258a3.61 3.61 0 01-1.59-.658l-1.375-1.01z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Upload returns a string representing an "upload" Octicon.
func Upload(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-upload%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M8.53 1.22a.75.75 0 00-1.06 0L3.72 4.97a.75.75 0 001.06 1.06l2.47-2.47v6.69a.75.75 0 001.5 0V3.56l2.47 2.47a.75.75 0 101.06-1.06L8.53 1.22zM3.75 13a.75.75 0 000 1.5h8.5a.75.75 0 000-1.5h-8.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-upload%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M4.75 22a.75.75 0 010-1.5h14.5a.75.75 0 010 1.5H4.75zm.22-13.53a.75.75 0 001.06 1.06L11 4.56v12.19a.75.75 0 001.5 0V4.56l4.97 4.97a.75.75 0 101.06-1.06l-6.25-6.25a.75.75 0 00-1.06 0L4.97 8.47z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Verified returns a string representing an "verified" Octicon.
func Verified(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-verified%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M9.585.52a2.678 2.678 0 00-3.17 0l-.928.68a1.178 1.178 0 01-.518.215L3.83 1.59a2.678 2.678 0 00-2.24 2.24l-.175 1.14a1.178 1.178 0 01-.215.518l-.68.928a2.678 2.678 0 000 3.17l.68.928c.113.153.186.33.215.518l.175 1.138a2.678 2.678 0 002.24 2.24l1.138.175c.187.029.365.102.518.215l.928.68a2.678 2.678 0 003.17 0l.928-.68a1.17 1.17 0 01.518-.215l1.138-.175a2.678 2.678 0 002.241-2.241l.175-1.138c.029-.187.102-.365.215-.518l.68-.928a2.678 2.678 0 000-3.17l-.68-.928a1.179 1.179 0 01-.215-.518L14.41 3.83a2.678 2.678 0 00-2.24-2.24l-1.138-.175a1.179 1.179 0 01-.518-.215L9.585.52zM7.303 1.728c.415-.305.98-.305 1.394 0l.928.68c.348.256.752.423 1.18.489l1.136.174c.51.078.909.478.987.987l.174 1.137c.066.427.233.831.489 1.18l.68.927c.305.415.305.98 0 1.394l-.68.928a2.678 2.678 0 00-.489 1.18l-.174 1.136a1.178 1.178 0 01-.987.987l-1.137.174a2.678 2.678 0 00-1.18.489l-.927.68c-.415.305-.98.305-1.394 0l-.928-.68a2.678 2.678 0 00-1.18-.489l-1.136-.174a1.178 1.178 0 01-.987-.987l-.174-1.137a2.678 2.678 0 00-.489-1.18l-.68-.927a1.178 1.178 0 010-1.394l.68-.928c.256-.348.423-.752.489-1.18l.174-1.136c.078-.51.478-.909.987-.987l1.137-.174a2.678 2.678 0 001.18-.489l.927-.68zM11.28 6.78a.75.75 0 00-1.06-1.06L7 8.94 5.78 7.72a.75.75 0 00-1.06 1.06l1.75 1.75a.75.75 0 001.06 0l3.75-3.75z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-verified%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M17.03 9.78a.75.75 0 00-1.06-1.06l-5.47 5.47-2.47-2.47a.75.75 0 00-1.06 1.06l3 3a.75.75 0 001.06 0l6-6z\"></path><path fill-rule=\"evenodd\" d=\"M14.136 1.2a3.61 3.61 0 00-4.272 0L8.489 2.21a2.11 2.11 0 01-.929.384l-1.686.259a3.61 3.61 0 00-3.021 3.02L2.594 7.56a2.11 2.11 0 01-.384.929L1.2 9.864a3.61 3.61 0 000 4.272l1.01 1.375c.2.274.333.593.384.929l.259 1.686a3.61 3.61 0 003.02 3.021l1.687.259c.336.051.655.183.929.384l1.375 1.01a3.61 3.61 0 004.272 0l1.375-1.01a2.11 2.11 0 01.929-.384l1.686-.259a3.61 3.61 0 003.021-3.02l.259-1.687a2.11 2.11 0 01.384-.929l1.01-1.375a3.61 3.61 0 000-4.272l-1.01-1.375a2.11 2.11 0 01-.384-.929l-.259-1.686a3.61 3.61 0 00-3.02-3.021l-1.687-.259a2.11 2.11 0 01-.929-.384L14.136 1.2zm-3.384 1.209a2.11 2.11 0 012.496 0l1.376 1.01a3.61 3.61 0 001.589.658l1.686.258a2.11 2.11 0 011.765 1.766l.26 1.686a3.61 3.61 0 00.657 1.59l1.01 1.375a2.11 2.11 0 010 2.496l-1.01 1.376a3.61 3.61 0 00-.658 1.589l-.258 1.686a2.11 2.11 0 01-1.766 1.765l-1.686.26a3.61 3.61 0 00-1.59.657l-1.375 1.01a2.11 2.11 0 01-2.496 0l-1.376-1.01a3.61 3.61 0 00-1.589-.658l-1.686-.258a2.11 2.11 0 01-1.766-1.766l-.258-1.686a3.61 3.61 0 00-.658-1.59l-1.01-1.375a2.11 2.11 0 010-2.496l1.01-1.376a3.61 3.61 0 00.658-1.589l.258-1.686a2.11 2.11 0 011.766-1.766l1.686-.258a3.61 3.61 0 001.59-.658l1.375-1.01z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Versions returns a string representing an "versions" Octicon.
func Versions(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-versions%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M7.75 14A1.75 1.75 0 016 12.25v-8.5C6 2.784 6.784 2 7.75 2h6.5c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0114.25 14h-6.5zm-.25-1.75c0 .138.112.25.25.25h6.5a.25.25 0 00.25-.25v-8.5a.25.25 0 00-.25-.25h-6.5a.25.25 0 00-.25.25v8.5zM4.9 3.508a.75.75 0 01-.274 1.025.25.25 0 00-.126.217v6.5a.25.25 0 00.126.217.75.75 0 01-.752 1.298A1.75 1.75 0 013 11.25v-6.5c0-.649.353-1.214.874-1.516a.75.75 0 011.025.274zM1.625 5.533a.75.75 0 10-.752-1.299A1.75 1.75 0 000 5.75v4.5c0 .649.353 1.214.874 1.515a.75.75 0 10.752-1.298.25.25 0 01-.126-.217v-4.5a.25.25 0 01.126-.217z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-versions%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10 22a2 2 0 01-2-2V4a2 2 0 012-2h11a2 2 0 012 2v16a2 2 0 01-2 2H10zm-.5-2a.5.5 0 00.5.5h11a.5.5 0 00.5-.5V4a.5.5 0 00-.5-.5H10a.5.5 0 00-.5.5v16zM6.17 4.165a.75.75 0 01-.335 1.006c-.228.114-.295.177-.315.201a.037.037 0 00-.008.016.387.387 0 00-.012.112v13c0 .07.008.102.012.112a.03.03 0 00.008.016c.02.024.087.087.315.201a.75.75 0 11-.67 1.342c-.272-.136-.58-.315-.81-.598C4.1 19.259 4 18.893 4 18.5v-13c0-.393.1-.759.355-1.073.23-.283.538-.462.81-.598a.75.75 0 011.006.336zM2.15 5.624a.75.75 0 01-.274 1.025c-.15.087-.257.17-.32.245C1.5 6.96 1.5 6.99 1.5 7v10c0 .01 0 .04.056.106.063.074.17.158.32.245a.75.75 0 11-.752 1.298C.73 18.421 0 17.907 0 17V7c0-.907.73-1.42 1.124-1.65a.75.75 0 011.025.274z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Video returns a string representing an "video" Octicon.
func Video(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-video%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 3.5a.25.25 0 00-.25.25v8.5c0 .138.112.25.25.25h12.5a.25.25 0 00.25-.25v-8.5a.25.25 0 00-.25-.25H1.75zM0 3.75C0 2.784.784 2 1.75 2h12.5c.966 0 1.75.784 1.75 1.75v8.5A1.75 1.75 0 0114.25 14H1.75A1.75 1.75 0 010 12.25v-8.5z\"></path><path d=\"M6 10.559V5.442a.25.25 0 01.379-.215l4.264 2.559a.25.25 0 010 .428l-4.264 2.559A.25.25 0 016 10.559z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-video%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.75 4.5a.25.25 0 00-.25.25v14.5c0 .138.112.25.25.25h20.5a.25.25 0 00.25-.25V4.75a.25.25 0 00-.25-.25H1.75zM0 4.75C0 3.784.784 3 1.75 3h20.5c.966 0 1.75.784 1.75 1.75v14.5A1.75 1.75 0 0122.25 21H1.75A1.75 1.75 0 010 19.25V4.75z\"></path><path d=\"M9 15.584V8.416a.5.5 0 01.77-.42l5.576 3.583a.5.5 0 010 .842L9.77 16.005a.5.5 0 01-.77-.42z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Webhook returns a string representing an "webhook" Octicon.
func Webhook(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-webhook%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path d=\"M5.5 4.25a2.25 2.25 0 014.5 0 .75.75 0 001.5 0 3.75 3.75 0 10-6.14 2.889l-2.272 4.258a.75.75 0 001.324.706L7 7.25a.75.75 0 00-.309-1.015A2.25 2.25 0 015.5 4.25z\"></path><path d=\"M7.364 3.607a.75.75 0 011.03.257l2.608 4.349a3.75 3.75 0 11-.628 6.785.75.75 0 01.752-1.299 2.25 2.25 0 10-.033-3.88.75.75 0 01-1.03-.256L7.107 4.636a.75.75 0 01.257-1.03z\"></path><path d=\"M2.9 8.776A.75.75 0 012.625 9.8 2.25 2.25 0 106 11.75a.75.75 0 01.75-.751h5.5a.75.75 0 010 1.5H7.425a3.751 3.751 0 11-5.55-3.998.75.75 0 011.024.274z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Workflow returns a string representing an "workflow" Octicon.
func Workflow(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-workflow%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M0 1.75C0 .784.784 0 1.75 0h3.5C6.216 0 7 .784 7 1.75v3.5A1.75 1.75 0 015.25 7H4v4a1 1 0 001 1h4v-1.25C9 9.784 9.784 9 10.75 9h3.5c.966 0 1.75.784 1.75 1.75v3.5A1.75 1.75 0 0114.25 16h-3.5A1.75 1.75 0 019 14.25v-.75H5A2.5 2.5 0 012.5 11V7h-.75A1.75 1.75 0 010 5.25v-3.5zm1.75-.25a.25.25 0 00-.25.25v3.5c0 .138.112.25.25.25h3.5a.25.25 0 00.25-.25v-3.5a.25.25 0 00-.25-.25h-3.5zm9 9a.25.25 0 00-.25.25v3.5c0 .138.112.25.25.25h3.5a.25.25 0 00.25-.25v-3.5a.25.25 0 00-.25-.25h-3.5z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-workflow%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1 3a2 2 0 012-2h6.5a2 2 0 012 2v6.5a2 2 0 01-2 2H7v4.063C7 16.355 7.644 17 8.438 17H12.5v-2.5a2 2 0 012-2H21a2 2 0 012 2V21a2 2 0 01-2 2h-6.5a2 2 0 01-2-2v-2.5H8.437A2.938 2.938 0 015.5 15.562V11.5H3a2 2 0 01-2-2V3zm2-.5a.5.5 0 00-.5.5v6.5a.5.5 0 00.5.5h6.5a.5.5 0 00.5-.5V3a.5.5 0 00-.5-.5H3zM14.5 14a.5.5 0 00-.5.5V21a.5.5 0 00.5.5H21a.5.5 0 00.5-.5v-6.5a.5.5 0 00-.5-.5h-6.5z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// X returns a string representing an "x" Octicon.
func X(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-x%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.72 3.72a.75.75 0 011.06 0L8 6.94l3.22-3.22a.75.75 0 111.06 1.06L9.06 8l3.22 3.22a.75.75 0 11-1.06 1.06L8 9.06l-3.22 3.22a.75.75 0 01-1.06-1.06L6.94 8 3.72 4.78a.75.75 0 010-1.06z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-x%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M5.72 5.72a.75.75 0 011.06 0L12 10.94l5.22-5.22a.75.75 0 111.06 1.06L13.06 12l5.22 5.22a.75.75 0 11-1.06 1.06L12 13.06l-5.22 5.22a.75.75 0 01-1.06-1.06L10.94 12 5.72 6.78a.75.75 0 010-1.06z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// XCircle returns a string representing an "x-circle" Octicon.
func XCircle(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-x-circle%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M3.404 12.596a6.5 6.5 0 119.192-9.192 6.5 6.5 0 01-9.192 9.192zM2.344 2.343a8 8 0 1011.313 11.314A8 8 0 002.343 2.343zM6.03 4.97a.75.75 0 00-1.06 1.06L6.94 8 4.97 9.97a.75.75 0 101.06 1.06L8 9.06l1.97 1.97a.75.75 0 101.06-1.06L9.06 8l1.97-1.97a.75.75 0 10-1.06-1.06L8 6.94 6.03 4.97z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-x-circle%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path d=\"M9.036 7.976a.75.75 0 00-1.06 1.06L10.939 12l-2.963 2.963a.75.75 0 101.06 1.06L12 13.06l2.963 2.964a.75.75 0 001.061-1.06L13.061 12l2.963-2.964a.75.75 0 10-1.06-1.06L12 10.939 9.036 7.976z\"></path><path fill-rule=\"evenodd\" d=\"M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11 11-4.925 11-11S18.075 1 12 1zM2.5 12a9.5 9.5 0 1119 0 9.5 9.5 0 01-19 0z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// XCircleFill returns a string representing an "x-circle-fill" Octicon.
func XCircleFill(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 12:
		return fmt.Sprintf("<svg class=\"octicon octicon-x-circle-fill%s\" height=\"12\" width=\"12\" viewbox=\"0 0 12 12\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1.757 10.243a6 6 0 118.486-8.486 6 6 0 01-8.486 8.486zM6 4.763l-2-2L2.763 4l2 2-2 2L4 9.237l2-2 2 2L9.237 8l-2-2 2-2L8 2.763l-2 2z\"></path></svg>", extraClasses), true
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-x-circle-fill%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M2.343 13.657A8 8 0 1113.657 2.343 8 8 0 012.343 13.657zM6.03 4.97a.75.75 0 00-1.06 1.06L6.94 8 4.97 9.97a.75.75 0 101.06 1.06L8 9.06l1.97 1.97a.75.75 0 101.06-1.06L9.06 8l1.97-1.97a.75.75 0 10-1.06-1.06L8 6.94 6.03 4.97z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-x-circle-fill%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11-4.925 11-11 11S1 18.075 1 12zm8.036-4.024a.75.75 0 00-1.06 1.06L10.939 12l-2.963 2.963a.75.75 0 101.06 1.06L12 13.06l2.963 2.964a.75.75 0 001.061-1.06L13.061 12l2.963-2.964a.75.75 0 10-1.06-1.06L12 10.939 9.036 7.976z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}

// Zap returns a string representing an "zap" Octicon.
func Zap(height int, rawExtraClasses ...string) (string, bool) {
	extraClasses := strings.Join(rawExtraClasses, " ")
	if extraClasses != "" {
		extraClasses = " " + extraClasses
	}
	switch height {
	case 16:
		return fmt.Sprintf("<svg class=\"octicon octicon-zap%s\" height=\"16\" width=\"16\" viewbox=\"0 0 16 16\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M10.561 1.5a.016.016 0 00-.01.004L3.286 8.571A.25.25 0 003.462 9H6.75a.75.75 0 01.694 1.034l-1.713 4.188 6.982-6.793A.25.25 0 0012.538 7H9.25a.75.75 0 01-.683-1.06l2.008-4.418.003-.006a.02.02 0 00-.004-.009.02.02 0 00-.006-.006L10.56 1.5zM9.504.43a1.516 1.516 0 012.437 1.713L10.415 5.5h2.123c1.57 0 2.346 1.909 1.22 3.004l-7.34 7.142a1.25 1.25 0 01-.871.354h-.302a1.25 1.25 0 01-1.157-1.723L5.633 10.5H3.462c-1.57 0-2.346-1.909-1.22-3.004L9.503.429z\"></path></svg>", extraClasses), true
	case 24:
		return fmt.Sprintf("<svg class=\"octicon octicon-zap%s\" height=\"24\" width=\"24\" viewbox=\"0 0 24 24\" aria-hidden=\"true\"><path fill-rule=\"evenodd\" d=\"M16.168 2.924L4.51 13.061a.25.25 0 00.164.439h5.45a.75.75 0 01.692 1.041l-2.559 6.066 11.215-9.668a.25.25 0 00-.164-.439H14a.75.75 0 01-.687-1.05l2.855-6.526zm-.452-1.595a1.341 1.341 0 012.109 1.55L15.147 9h4.161c1.623 0 2.372 2.016 1.143 3.075L8.102 22.721a1.149 1.149 0 01-1.81-1.317L8.996 15H4.674c-1.619 0-2.37-2.008-1.148-3.07l12.19-10.6z\"></path></svg>", extraClasses), true
	default:
		return "", false
	}
}
