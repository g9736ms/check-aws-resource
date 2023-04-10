package env_reader

import (
	"bufio"
	"os"
	"strings"
)

func EnvReader(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	env := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	//return env, nil
	return env, err

}

// 암호화 되서 올린다.
// ec2 에 가져올수 있다 . get ec2 tags <-

// SetEnvFromFile sets environment variables from a file.
func SetEnv(filename string) error {
	env, err := EnvReader(filename)
	if err != nil {
		return err
	}

	for k, v := range env {
		err = os.Setenv(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetEnv returns the value of the specified environment variable, or an empty string if it is not set.
func GetEnv(key string) string {
	return os.Getenv(key)
}

// 와퍼클래스 < 찾아 봐야한다
/*
func main() {
	// Load environment variables from file
	err := env_reader.SetEnv("/tmp/test.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 값을 가져와 사용
	dbHost := env_reader.GetEnv("DB_HOST")
	fmt.Println("DB_HOST:", dbHost)
}

*/
