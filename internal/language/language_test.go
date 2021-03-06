package language

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestLanguage(t *testing.T) {
	var l Language
	assert.Equal(t, Unset, l)

	assert.Empty(t, Bash.Executable().Name())
	assert.True(t, Bash.Executable().Builtin())

	assert.NotEmpty(t, Python3.Executable().Name())
	assert.False(t, Python3.Executable().Builtin())

	assert.Equal(t, "#!/usr/bin/env perl\n", Perl.Header())
}

func TestMakeLanguage(t *testing.T) {
	assert.Equal(t, Python3, MakeByName("python3"), "python3")
	assert.Equal(t, Unknown, MakeByName("python4"), "unknown language")
	assert.Equal(t, Unset, MakeByName(""), "unset language")
}

func TestUnmarshal(t *testing.T) {
	var l Language

	err := yaml.Unmarshal([]byte("junk"), &l)
	assert.Error(t, err, "fail due to bad yaml input")
	assert.Equal(t, Unset, l)

	err = yaml.Unmarshal([]byte("python3"), &l)
	assert.NoError(t, err, "successfully unmarshal 'python3'")
	assert.Equal(t, Python3, l)

	err = yaml.Unmarshal([]byte("bash"), &l)
	assert.NoError(t, err, "successfully unmarshal 'bash'")
	assert.Equal(t, Bash, l)

	err = yaml.Unmarshal([]byte("unknown"), &l)
	assert.Error(t, err, "not successfully unmarshal 'unknown'")
}

func TestMarshal(t *testing.T) {
	l := Python3
	bs, err := yaml.Marshal(l)
	require.NoError(t, err)
	assert.Contains(t, string(bs), "python")

	l = Batch
	bs, err = yaml.Marshal(&l)
	require.NoError(t, err)
	assert.Contains(t, string(bs), "batch")
	assert.Empty(t, l.Header())

}

func TestMakeLanguageByShell(t *testing.T) {
	assert.Equal(t, Batch, MakeByShell("cmd.exe"), "strings with 'cmd' return batch")
	assert.Equal(t, Bash, MakeByShell("anything_else"), "anything else returns bash")
}

func TestRecognized(t *testing.T) {
	langs := Recognized()
	for _, l := range langs {
		assert.NotEqual(t, l, Unset, "not unset")
		assert.NotEqual(t, l, Unknown, "not unknown")
	}
}

func TestSupported(t *testing.T) {
	var l Supported
	assert.Equal(t, Unset, l.Language)
}

func TestSupportedUnmarshal(t *testing.T) {
	var l Supported

	err := yaml.Unmarshal([]byte("junk"), &l)
	assert.Error(t, err, "fail due to bad yaml input")
	assert.Equal(t, Unset, l.Language)

	err = yaml.Unmarshal([]byte("python3"), &l)
	assert.NoError(t, err, "successfully unmarshal 'python3'")
	assert.Equal(t, Python3, l.Language)

	err = yaml.Unmarshal([]byte("bash"), &l)
	assert.Error(t, err, "not successfully unmarshal 'bash'")
}

func TestSupportedMarshal(t *testing.T) {
	l := Supported{Python3}
	bs, err := yaml.Marshal(l)
	require.NoError(t, err)
	assert.Contains(t, string(bs), "python")

	l = Supported{Batch}
	bs, err = yaml.Marshal(&l)
	require.NoError(t, err)
	assert.Contains(t, string(bs), "batch")
	assert.Empty(t, l.Header())

}

func TestRecognizedSupporteds(t *testing.T) {
	langs := RecognizedSupporteds()
	for _, l := range langs {
		assert.NotEqual(t, l.Language, Unset, "not unset")
		assert.NotEqual(t, l.Language, Unknown, "not unknown")
		assert.False(t, l.Executable().Builtin())
		assert.NotEmpty(t, l.Executable().Name())
	}
}
