# Docker MySQL 環境構築ガイド

## 1. Dockerfile
最小構成のDockerfile例:

```dockerfile
FROM mysql:8.0

# 文字コードの設定
ENV LANG="ja_JP.UTF-8"
ENV LC_ALL="ja_JP.UTF-8"
ENV LANGUAGE="ja_JP.UTF-8"

# タイムゾーンの設定
ENV TZ="Asia/Tokyo"

# MySQLの環境変数設定
ENV MYSQL_USER=jboy
ENV MYSQL_PASSWORD=1234qw
ENV MYSQL_ROOT_PASSWORD=root1234
ENV MYSQL_DATABASE=myapp
```

## 2. 基本的な使い方

### イメージのビルド
```bash
docker build -t mysql-custom .
```

### コンテナの起動
```bash
docker run -d -p 3306:3306 --name mysql-container mysql-custom
```

### MySQLへの接続
- rootユーザーとして接続:
```bash
docker exec -it mysql-container mysql -u root -p
# パスワード: root1234
```

- 一般ユーザーとして接続:
```bash
docker exec -it mysql-container mysql -u jboy -p
# パスワード: 1234qw
```

## 3. 接続情報
- ホスト: localhost
- ポート: 3306
- データベース: myapp
- ユーザー名: jboy
- パスワード: 1234qw

## 4. 運用コマンド

### コンテナの管理
```bash
# コンテナの停止
docker stop mysql-container

# コンテナの削除
docker rm mysql-container

# コンテナの状態確認
docker ps

# コンテナのログ確認
docker logs mysql-container
```

### 環境の再構築
```bash
# 1. 既存のコンテナを停止・削除
docker stop mysql-container
docker rm mysql-container

# 2. イメージの再ビルド
docker build -t mysql-custom .

# 3. 新しいコンテナの起動
docker run -d -p 3306:3306 --name mysql-container mysql-custom
```