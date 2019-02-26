package constants

// VersionNumber holds the current version of our cli
const VersionNumber = "0.2.2"

// LibraryName contains the main name of this library
const LibraryName = "cli"

// LibraryNamespace is the namespace that the library belongs to
const LibraryNamespace = "github.com/ActiveState/"

// CommandName holds the name of our command
const CommandName = "state"

// ConfigNamespace holds the appdata folder name under which we store our config
const ConfigNamespace = "activestate"

// ConfigName is used to inform viper and our config lib about the name of the config file
const ConfigName = "activestate"

// ConfigFileName is effectively the same as ConfigName, but includes our preferred extension
const ConfigFileName = ConfigName + ".yaml"

// ConfigFileType is our preferred file type for our config file, this must match ConfigFileName
const ConfigFileType = "yaml"

// EnvironmentEnvVarName is the name of the environment variable that specifies the current environment (dev, qa, prod, etc.)
const EnvironmentEnvVarName = "ACTIVESTATE_ENVIRONMENT"

// ProjectEnvVarName is the name of the environment variable that specifies the path of the activestate.yaml config file.
const ProjectEnvVarName = "ACTIVESTATE_PROJECT"

// ActivatedStateEnvVarName is the name of the environment variable that is set when in an activated state, its value will be the path of the project
const ActivatedStateEnvVarName = "ACTIVESTATE_ACTIVATED"

// ForwardedStateEnvVarName is the name of the environment variable that is set when in an activated state, its value will be the path of the project
const ForwardedStateEnvVarName = "ACTIVESTATE_FORWARDED"

// APIUpdateURL is the URL for our update server
const APIUpdateURL = "https://s3.ca-central-1.amazonaws.com/cli-update/update/"

// APIArtifactURL is the URL for downloading artifacts
const APIArtifactURL = "https://s3.ca-central-1.amazonaws.com/cli-artifacts/"

// ArtifactFile is the name of the artifact json file contained within artifacts
const ArtifactFile = "artifact.json"

// UpdateStorageDir is the directory where updates will be stored
const UpdateStorageDir = "update/"

// DefaultNamespaceDomain is the domain used when no namespace is given and one has to be constructed
const DefaultNamespaceDomain = "github.com"

// AnalyticsTrackingID is our Google Analytics tracking ID
const AnalyticsTrackingID = "UA-118120158-1"

// APITokenName is the name we give our api token
const APITokenName = "activestate-platform-cli"

// KeypairLocalFileName is the name of the file (sans extension) that will hold the user's unencrypted
// private key in their config dir.
const KeypairLocalFileName = "private"

// DefaultRSABitLength represents the default RSA bit-length that will be assumed when
// generating new Keypairs.
const DefaultRSABitLength int = 4096

// ExpanderMaxDepth defines the maximum depth to fully expand a given value.
const ExpanderMaxDepth = int(10)

// StableBranch is the branch mapped to stable builds
const StableBranch = "stable"

// UnstableBranch is the branch used for unstable builds
const UnstableBranch = "unstable"

// ExperimentalBranch is the branch used for experimental builds
const ExperimentalBranch = "master"

// PlatformAPIPath is the api path used for the platform api
const PlatformAPIPath = "/api/v1"

// SecretsAPIPath is the api path used for the secrets api
const SecretsAPIPath = "/api/secrets/v1"

// PlatformURLProd is the host used for platform api calls when on production
const PlatformURLProd = "https://platform.activestate.com" + PlatformAPIPath

// SecretsURLProd is the host used for secrets api calls when on production
const SecretsURLProd = "https://platform.activestate.com" + SecretsAPIPath

// PlatformURLStage is the host used for platform api calls when on staging
const PlatformURLStage = "https://staging.activestate.build" + PlatformAPIPath

// SecretsURLStage is the host used for secrets api calls when on staging
const SecretsURLStage = "https://staging.activestate.com" + SecretsAPIPath

// PlatformURLDev is the host used for platform api calls when on staging
const PlatformURLDev = PlatformURLStage

// SecretsURLDev is the host used for secrets api calls when on dev
const SecretsURLDev = "http://localhost:8080" + SecretsAPIPath
