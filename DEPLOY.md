# Deployment Guide

## 🚀 Quick Deploy Options

### 1. Railway.app (Recommended - Free Tier Available)

1. **Fork this repository** to your GitHub account

2. **Connect to Railway**:
   - Go to [Railway.app](https://railway.app)
   - Sign up with GitHub
   - Click "New Project" → "Deploy from GitHub repo"
   - Select your forked regret-box repository

3. **Add Redis**:
   - In your Railway project dashboard
   - Click "New" → "Database" → "Add Redis"

4. **Set Environment Variables**:
   - Go to your app service settings
   - Add environment variables:
     - `REDIS_URL`: `redis://redis.railway.internal:6379`
     - `PORT`: `8080`

5. **Deploy**: Railway will automatically build and deploy!

### 2. DigitalOcean App Platform

1. **Create new app**:
   - Go to [DigitalOcean Apps](https://cloud.digitalocean.com/apps)
   - Connect your GitHub repo

2. **Add Redis database**:
   - Add a Dev Database (Redis)

3. **Configure environment**:
   - Set `REDIS_URL` to your Redis connection string
   - Set `PORT` to `8080`

### 3. Heroku

```bash
# Install Heroku CLI, then:
heroku create your-regretbox-name
heroku addons:create heroku-redis:mini
git push heroku main
```

### 4. VPS with Docker

```bash
# On your VPS:
git clone <your-repo>
cd regret-box
docker-compose -f docker-compose.prod.yml up -d
```

## 💡 Tips

- **Custom Domain**: Most platforms allow adding custom domains
- **SSL**: Automatic HTTPS on most cloud platforms
- **Scaling**: All platforms support easy scaling
- **Monitoring**: Enable logging and monitoring in production

## 🔧 Environment Variables

Required for all deployments:
- `PORT`: Port number (usually 8080)
- `REDIS_URL`: Redis connection string

## 📋 Pre-deployment Checklist

- [ ] Repository is public or accessible to deployment platform
- [ ] Go 1.22+ is available (handled automatically by Docker/cloud platforms)
- [ ] Environment variables are set correctly
- [ ] Redis database is provisioned
- [ ] Build commands are configured (usually automatic)
- [ ] Domain is configured (optional)

## 🎯 Recommended: Railway.app

Railway.app is the easiest option because:
- ✅ Free tier available
- ✅ Automatic Redis setup
- ✅ GitHub integration
- ✅ Automatic builds and deployments
- ✅ Easy environment variable management
- ✅ Built-in monitoring and logs
