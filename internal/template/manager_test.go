package template

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arch-err/tmux-hive/internal/config"
)

func TestGetTemplateDir(t *testing.T) {
	// Save original env
	originalXDG := os.Getenv("XDG_DATA_HOME")
	defer os.Setenv("XDG_DATA_HOME", originalXDG)

	t.Run("with XDG_DATA_HOME set", func(t *testing.T) {
		os.Setenv("XDG_DATA_HOME", "/tmp/data")
		dir, err := GetTemplateDir()
		if err != nil {
			t.Fatalf("GetTemplateDir() error = %v", err)
		}

		expected := "/tmp/data/hive/templates"
		if dir != expected {
			t.Errorf("GetTemplateDir() = %q, want %q", dir, expected)
		}
	})

	t.Run("without XDG_DATA_HOME", func(t *testing.T) {
		os.Unsetenv("XDG_DATA_HOME")
		dir, err := GetTemplateDir()
		if err != nil {
			t.Fatalf("GetTemplateDir() error = %v", err)
		}

		if !filepath.IsAbs(dir) {
			t.Errorf("GetTemplateDir() should return absolute path, got %q", dir)
		}

		if !filepath.HasPrefix(dir, filepath.Join(os.Getenv("HOME"), ".local", "share")) {
			t.Errorf("GetTemplateDir() should use ~/.local/share, got %q", dir)
		}
	})
}

func TestListTemplates(t *testing.T) {
	templates, err := ListTemplates()
	if err != nil {
		t.Fatalf("ListTemplates() error = %v", err)
	}

	// Check that built-in templates are present
	builtins := []string{"basic", "dev", "ctf", "web", "blank", "minimal"}
	for _, name := range builtins {
		found := false
		for _, tmpl := range templates {
			if tmpl == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("ListTemplates() missing built-in template %q", name)
		}
	}
}

func TestReadTemplate(t *testing.T) {
	t.Run("read built-in template", func(t *testing.T) {
		cfg, err := ReadTemplate("basic")
		if err != nil {
			t.Fatalf("ReadTemplate(basic) error = %v", err)
		}

		if cfg == nil {
			t.Fatal("ReadTemplate(basic) returned nil")
		}

		if cfg.Session.Name != "basic" {
			t.Errorf("basic template session name = %q, want %q", cfg.Session.Name, "basic")
		}
	})

	t.Run("read all built-in templates", func(t *testing.T) {
		templates := []string{"basic", "dev", "ctf", "web", "blank", "minimal"}
		for _, name := range templates {
			cfg, err := ReadTemplate(name)
			if err != nil {
				t.Errorf("ReadTemplate(%q) error = %v", name, err)
				continue
			}
			if cfg == nil {
				t.Errorf("ReadTemplate(%q) returned nil", name)
			}
		}
	})

	t.Run("read nonexistent template", func(t *testing.T) {
		_, err := ReadTemplate("nonexistent-template")
		if err == nil {
			t.Error("ReadTemplate(nonexistent) should return error")
		}
	})
}

func TestSaveAndDeleteTemplate(t *testing.T) {
	// Use temporary directory
	tmpDir := t.TempDir()
	os.Setenv("XDG_DATA_HOME", tmpDir)
	defer os.Unsetenv("XDG_DATA_HOME")

	templateName := "test-template"
	cfg := &config.Config{
		Session: config.SessionConfig{
			Name: templateName,
		},
		Windows: []config.WindowConfig{
			{
				Name: "main",
				Panes: []config.PaneConfig{
					{Cmd: "echo test"},
				},
			},
		},
	}

	// Test SaveTemplate
	err := SaveTemplate(templateName, cfg)
	if err != nil {
		t.Fatalf("SaveTemplate() error = %v", err)
	}

	// Verify template was saved
	templatePath := filepath.Join(tmpDir, "hive", "templates", templateName+".yaml")
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		t.Errorf("Template file not created at %q", templatePath)
	}

	// Test reading saved template
	loaded, err := ReadTemplate(templateName)
	if err != nil {
		t.Fatalf("ReadTemplate() after save error = %v", err)
	}

	if loaded.Session.Name != cfg.Session.Name {
		t.Errorf("Loaded template session name = %q, want %q", loaded.Session.Name, cfg.Session.Name)
	}

	// Test DeleteTemplate
	err = DeleteTemplate(templateName)
	if err != nil {
		t.Fatalf("DeleteTemplate() error = %v", err)
	}

	// Verify template was deleted
	if _, err := os.Stat(templatePath); !os.IsNotExist(err) {
		t.Error("Template file still exists after delete")
	}
}

func TestDeleteBuiltinTemplate(t *testing.T) {
	err := DeleteTemplate("basic")
	if err == nil {
		t.Error("DeleteTemplate(basic) should return error for built-in template")
	}
}

func TestEnsureTemplateDir(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("XDG_DATA_HOME", tmpDir)
	defer os.Unsetenv("XDG_DATA_HOME")

	err := EnsureTemplateDir()
	if err != nil {
		t.Fatalf("EnsureTemplateDir() error = %v", err)
	}

	templateDir := filepath.Join(tmpDir, "hive", "templates")
	info, err := os.Stat(templateDir)
	if err != nil {
		t.Fatalf("Template directory not created: %v", err)
	}

	if !info.IsDir() {
		t.Error("Template path is not a directory")
	}
}
