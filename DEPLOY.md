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
   - Note the service name (usually "Redis")

4. **Set Environment Variables**:
   - Go to your app service settings → Variables tab
   - Add environment variables:
     - `REDIS_URL`: `${{Redis.REDIS_URL}}` (this references your Redis service)
     - `PORT`: `8080`
   
   **Important**: Use the Railway template variable `${{Redis.REDIS_URL}}` to automatically reference your Redis service. Replace "Redis" with your actual Redis service name if different.

5. **Deploy**: Railway will automatically build and deploy!

#### Railway.app Troubleshooting:

**If you see connection refused errors:**
1. Check that your Redis service is running in the Railway dashboard
2. Verify the environment variable is set to `${{Redis.REDIS_URL}}` (with your actual service name)
3. Make sure both services are in the same project
4. Check the logs for "Using Redis URL:" to see what URL is being used

**Service Reference Format:**
- If your Redis service is named "Redis": `${{Redis.REDIS_URL}}`
- If your Redis service is named "redis-db": `${{redis-db.REDIS_URL}}`
- The format is: `${{SERVICE_NAME.REDIS_URL}}`

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

## 🔧 Troubleshooting

### Redis Authentication Errors

If you see `NOAUTH Authentication required` errors:

1. **Check Redis URL format**: Make sure you're using the full Redis URL with credentials:
   ```
   redis://:password@host:port
   # or
   redis://username:password@host:port
   ```

2. **Railway.app**: Get the correct Redis URL from:
   - Go to your Redis service
   - Click "Connect" tab
   - Copy the full connection string
   - Set it as `REDIS_URL` environment variable

3. **Local development**: If using local Redis with password:
   ```bash
   export REDIS_PASSWORD=your_password
   export REDIS_URL=localhost:6379
   ```

### Build Failures

- **Go version issues**: Ensure Go 1.22+ is used (automatic in Docker/cloud)
- **Missing dependencies**: Run `go mod tidy` locally to verify

### Connection Issues

- **Port conflicts**: Change port in environment variables if needed
- **Firewall**: Ensure ports 8080 and 6379 are open for local development
