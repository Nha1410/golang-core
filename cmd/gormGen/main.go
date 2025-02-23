package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	gormGen()
	excludeMigration()
	cleanup()
	// generateHook()
	// generateOptimisticLock()
}

const (
	generatedPath      = "./db/generated"
	generatedModelPath = "./db/generated/model"
	tmpPath            = "./db/generated/tmp"
	modelDir           = "model"
	queryDir           = "query"
	migrateFilename    = "__migrations.gen.go"
	genFilename        = "gen.go"
	hookFilename       = "hook.go"
)

func gormGen() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "db/generated/tmp/query",
		ModelPkgPath:      "db/generated/tmp/model",
		WithUnitTest:      false,
		FieldNullable:     true,
		FieldCoverable:    true, // defaultが設定されている場合,zero値を代入できなくする設定
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode: gen.WithoutContext |
			gen.WithDefaultQuery |
			gen.WithQueryInterface,
	})
	dsn := "host=fiber_postgres user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	// os.Getenv("DB_HOST"),
	// os.Getenv("DB_USER"),
	// os.Getenv("DB_PASSWORD"),
	// os.Getenv("DB_NAME"),
	// os.Getenv("DB_PORT")

	gormDb, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	tableList, err := gormDb.Migrator().GetTables()
	if err != nil {
		panic(fmt.Errorf("get all tables fail: %w", err))
	}

	modelNames := map[string]string{}
	for _, tableName := range tableList {
		modelNames[tableName] = gormDb.NamingStrategy.SchemaName(tableName)
	}

	db, err := gormDb.DB()
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}

	gormDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			IdentifierMaxLength: 64, // Default Identifier length is 64
			SingularTable:       true,
		},
	})

	if err != nil {
		panic(err)
	}

	g.UseDB(gormDb)

	// Generate models
	tables := make([]interface{}, len(tableList))
	for i, tableName := range tableList {
		tables[i] = g.GenerateModelAs(tableName, modelNames[tableName])
	}

	// Apply basic settings to models and generate the code
	g.ApplyBasic(tables...)
	g.Execute()
}

func excludeMigration() {
	for _, v := range []string{path.Join(tmpPath, modelDir, migrateFilename), path.Join(tmpPath, queryDir, migrateFilename)} {
		//ファイル存在チェック
		if _, err := os.Stat(v); os.IsNotExist(err) {
			continue
		}
		err := os.Remove(v)
		if err != nil {
			panic(err)
		}
	}
	genFilePath := path.Join(tmpPath, queryDir, genFilename)
	file, err := os.Open(genFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// "migration"または"Migration"を含まない行だけを新しいスライスに追加
		if !strings.Contains(strings.ToLower(line), "migration") {
			lines = append(lines, line)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// 新しいスライスの内容を元のファイルに書き戻す
	err = os.WriteFile(genFilePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}

func cleanup() {
	dirs := []string{modelDir, queryDir}
	for _, v := range dirs {
		base := fmt.Sprintf("%s/%s", generatedPath, v)
		tmp := fmt.Sprintf("%s/%s", tmpPath, v)
		remove(base, tmp)
		upload(base, tmp)
	}
	// tmpの削除
	err := os.Remove(tmpPath)
	if err != nil {
		panic(err)
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func remove(base, tmp string) {
	files, _ := os.ReadDir(base)
	for _, f := range files {
		filename := f.Name()
		filePath := path.Join(base, filename)
		fmt.Println(filePath)
		if filename == "hook.go" {
			continue
		}
		if !exists(path.Join(tmp, filename)) {
			err := os.Remove(filePath)
			if err != nil {
				panic(err)
			}
		}
	}
}

func upload(base, tmp string) {
	// 元のディレクトリを削除
	err := os.RemoveAll(base)
	if err != nil {
		panic(err)
	}

	// 新しいディレクトリを元のディレクトリの場所に移動
	err = os.Rename(tmp, base)
	if err != nil {
		panic(err)
	}

	// baseのディレクトリのファイルにある文字列の置換
	// /generated/tmp/ -> /generated/
	err = filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			read, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}

			newContents := strings.Replace(string(read), "/generated/tmp/", "/generated/", -1)

			err = os.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}
