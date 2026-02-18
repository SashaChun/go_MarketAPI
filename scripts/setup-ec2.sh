#!/bin/bash

set -e

echo "🚀 Setting up EC2 instance for deployment..."

sudo yum update -y

echo "📦 Installing Docker..."
sudo yum install -y docker
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker ec2-user

echo "📦 Installing Docker Compose..."
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

echo "📦 Installing Git..."
sudo yum install -y git

echo "🔑 Setting up project directory..."
cd ~
git clone https://github.com/YOUR_USERNAME/untitled9.git
cd untitled9

echo "📝 Creating .env file..."
cat > .env << 'ENVFILE'
PORT=5432
HOST=postgres
DATABASE_URL=shopdb
DB_USER=shop
PASSWORD=shop_pass
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
ENVFILE

echo "✅ EC2 setup complete!"
echo "⚠️  Don't forget to:"
echo "   1. Update .env with production values"
echo "   2. Configure security groups (ports 22, 8080, 5432)"
echo "   3. Add GitHub secrets: EC2_SSH_KEY, EC2_HOST, EC2_USER"
