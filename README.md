# 💌 RegretBox

A beautiful, anonymous confession platform where regrets vanish after 24 hours. Share your deepest regrets anonymously and connect with others through shared human experiences.

## ✨ Features

- **Anonymous Confessions**: Share your regrets completely anonymously
- **24-Hour Expiry**: All regrets automatically disappear after 24 hours
- **Like System**: Show empathy by liking regrets that resonate with you
- **Popular Page**: View the top 5 most liked regrets
- **Random Discovery**: Discover random regrets from others
- **Beautiful UI**: Modern, responsive design with smooth animations
- **Mobile Friendly**: Optimized for all device sizes
- **Real-time AJAX**: Like without page refreshes

## 🚀 Quick Start with Docker (Recommended)

### 1. Clone and Deploy
```bash
git clone <your-repo>
cd regret-box
./deploy.sh
```

### 2. Visit the App
```
http://localhost:8080
```

### 3. Check Status
```bash
docker-compose ps
docker-compose logs -f
```

## 🛠️ Manual Setup

### Prerequisites
- Go 1.21+
- Redis server

### Installation

1. **Install Redis**:
```bash
# Ubuntu/Debian
sudo apt install redis-server

# macOS
brew install redis

# Start Redis
redis-server
```

2. **Install dependencies**:
```bash
go mod tidy
```

3. **Run the application**:
```bash
go run main.go
```

## 🌐 Deployment Options

### 1. Railway.app (Easiest)
1. Connect your GitHub repo to Railway
2. Add Redis service
3. Set environment variables:
   - `PORT=8080`
   - `REDIS_URL=redis://redis:6379`
4. Deploy automatically

### 2. DigitalOcean App Platform
1. Create new app from GitHub repo
2. Add Redis database addon
3. Set environment variables
4. Deploy

### 3. VPS/Cloud Server
```bash
# Clone repo on server
git clone <your-repo>
cd regret-box

# Production deployment
docker-compose -f docker-compose.prod.yml up -d
```

### 4. Heroku
```bash
heroku create your-regretbox
heroku addons:create heroku-redis:mini
git push heroku main
```

## � Project Structure

```
regret-box/
├── main.go                 # Go backend server
├── go.mod                  # Go module dependencies
├── templates/              # HTML templates
│   ├── index.html         # Homepage
│   ├── regret.html        # Individual regret view
│   └── popular.html       # Popular regrets page
├── static/                # Static assets
│   ├── style.css         # Main stylesheet
│   └── script.js         # Frontend JavaScript
├── Dockerfile            # Container definition
├── docker-compose.yml    # Local development
├── docker-compose.prod.yml # Production deployment
├── deploy.sh             # Deployment script
└── README.md             # This file
```

## 🔧 Configuration

Environment variables:
- `PORT` - Server port (default: 8080)
- `REDIS_URL` - Redis connection string (default: localhost:6379)

## 📊 Monitoring & Management

### Check container status
```bash
docker-compose ps
```

### View logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f app
docker-compose logs -f redis
```

### Restart services
```bash
docker-compose restart
```

### Update deployment
```bash
git pull
./deploy.sh
```

### Database backup
```bash
docker-compose exec redis redis-cli BGSAVE
```

## 🎨 Features in Detail

### Anonymous Sharing
- No registration required
- No tracking of personal information
- IP-based like limiting (one like per regret per IP)

### Auto-Expiry System
- All regrets expire after exactly 24 hours
- Automatic cleanup using Redis TTL
- Like counts also expire with regrets

### Interactive Like System
- AJAX-based liking without page refreshes
- Real-time count updates with animations
- Visual feedback and notifications
- One like per IP address per regret

### Popular Regrets
- Top 5 most liked regrets
- Compact one-line layout
- Click to view full regret
- Real-time popularity ranking

### Modern UI/UX
- Gradient backgrounds with glass-morphism effects
- Smooth animations and micro-interactions
- Toast notifications for user feedback
- Dark mode support
- Fully responsive design

## � Security & Privacy

- **No personal data storage**: Only regret text and anonymous likes
- **IP-based rate limiting**: Prevents spam and abuse
- **Input sanitization**: All user input is properly escaped
- **Automatic data expiry**: No permanent data storage
- **No tracking**: No analytics or user tracking

## 🛠️ Development

### Local development
```bash
# Start Redis
docker run -d -p 6379:6379 redis:alpine

# Run app
go run main.go
```

### Hot reload (optional)
```bash
# Install air for hot reload
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Customization
- **Colors**: Edit CSS variables in `static/style.css`
- **Expiry time**: Change `24*time.Hour` in `main.go`
- **Popular count**: Modify limit in `popularHandler()`
- **Character limit**: Update JavaScript validation

## 🐛 Troubleshooting

### Redis connection issues
```bash
# Check Redis is running
docker-compose exec redis redis-cli ping
# Should return: PONG
```

### Port already in use
```bash
# Change port in docker-compose.yml
ports:
  - "8081:8080"  # Use port 8081 instead
```

### Build issues
```bash
# Clean rebuild
docker-compose down -v
docker-compose build --no-cache
docker-compose up -d
```

### Permission issues
```bash
chmod +x deploy.sh
```

## 💰 Hosting Costs

### Free Options:
- **Railway.app**: Free tier with limitations
- **Heroku**: Free tier (with sleep)

### Paid Options:
- **DigitalOcean App Platform**: ~$12/month (app + Redis)
- **VPS (DigitalOcean Droplet)**: ~$6/month
- **AWS/GCP/Azure**: Variable pricing

## � Quick Commands

```bash
# Deploy locally
./deploy.sh

# Deploy to production
docker-compose -f docker-compose.prod.yml up -d

# Check logs
docker-compose logs -f

# Stop everything
docker-compose down

# Clean rebuild
docker-compose down -v && docker-compose build --no-cache && docker-compose up -d
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

MIT License - feel free to use this project for your own purposes.

---

**Built with ❤️ for anonymous human connection**

*"Everyone has regrets. Here, they're shared, acknowledged, and then they fade away."*
