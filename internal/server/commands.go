package server

type Command string

const (
	CommandAbort Command             = "ABOR"
	CommandAccount Command           = "ACCT"
	CommandAuthData Command          = "ADAT"
	CommandAllo Command              = "ALLO"
	CommandAppend Command            = "APPE"
	CommandAuth Command              = "AUTH"
	CommandAvail Command             = "AVBL"
	CommandClear Command             = "CCC"
	CommandChangeParent Command      = "CDUP"
	CommandConf Command              = "CONF"
	CommandCsId Command              = "CSID"
	CommandChangeDir Command         = "CWD"
	CommandDelete Command            = "DELE"
	CommandDirSize Command           = "DSIZ"
	CommandPrivProtected Command     = "ENC"
	CommandExtAddrPort Command       = "EPRT"
	CommandExtPassMode Command       = "EPSV"
	CommandFeatLis Command           = "FEAT"
	CommandHelp Command              = "HELP"
	CommandHost Command              = "HOST"
	CommandLang Command              = "LANG"
	CommandList Command              = "LIST"
	CommandLongAddrPort Command      = "LPRT"
	CommandLongPassMode Command      = "LPSV"
	CommandLastModTime Command       = "MDTM"
	CommandModCreatTime Command      = "MFCT"
	CommandModFact Command           = "MFF"
	CommandModLastModTime Command    = "MFMT"
	CommandInteProtect Command       = "MIC"
	CommandMakeDir Command           = "MKD"
	CommandListDir Command           = "MLSD"
	CommandObjData Command           = "MLST"
	CommandMode Command              = "MODE"
	CommandFileNames Command         = "NLST"
	CommandNoOp Command              = "NOOP"
	CommandOptions Command           = "OPTS"
	CommandPassword Command          = "PASS"
	CommandPassive Command           = "PASV"
	CommandBufSizeProt Command       = "PBSZ"
	CommandPort Command              = "PORT"
	CommandDataChanProtLvl Command   = "PROT"
	CommandPrintDir Command          = "PWD"
	CommandQuit Command              = "QUIT"
	CommandReinit Command            = "REIN"
	CommandRestart Command           = "REST"
	CommandRetrieve Command          = "RETR"
	CommandRemoveDir Command         = "RMD"
	CommandRemoveDirTree Command     = "RMDA"
	CommandRenameFrom Command        = "RNFR"
	CommandRenameTo Command          = "RNTO"
	CommandSite Command              = "SITE"
	CommandFileSize Command          = "SIZE"
	CommandMountFile Command         = "SMNT"
	CommandSinglePortPassive Command = "SPSV"
	CommandServerStatus Command      = "STAT"
	CommandAcceptAndStore Command    = "STOR"
	CommandStoreFile Command         = "STOU"
	CommandFileStruct Command        = "STRU"
	CommandSystemType Command        = "SYST"
	CommandThumbnail Command         = "THMB"
	CommandType Command              = "TYPE"
	CommandUser Command              = "USER"
	CommandChangeToParentDir Command = "XCUP"
	CommandMakeADir Command          = "XMKD"
	CommandPrintCurDir Command       = "XPWD"
	CommandRemoveTheDir Command      = "XRMD"
	CommandSendMail Command          = "XSEM"
	CommandSendTerm Command          = "XSEN"
)