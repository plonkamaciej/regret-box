// Enhanced RegretBox JavaScript
document.addEventListener('DOMContentLoaded', function() {
    // Character counter for textarea
    const textarea = document.querySelector('textarea[name="regret"]');
    if (textarea) {
        const maxLength = 500;
        
        // Create character counter
        const counter = document.createElement('div');
        counter.className = 'char-counter';
        counter.style.cssText = `
            text-align: right;
            font-size: 0.9rem;
            color: #666;
            margin-top: 0.5rem;
        `;
        textarea.parentNode.appendChild(counter);
        
        function updateCounter() {
            const remaining = maxLength - textarea.value.length;
            counter.textContent = `${textarea.value.length}/${maxLength} characters`;
            
            if (remaining < 50) {
                counter.style.color = '#ff6b6b';
            } else if (remaining < 100) {
                counter.style.color = '#ffa500';
            } else {
                counter.style.color = '#666';
            }
        }
        
        textarea.addEventListener('input', updateCounter);
        textarea.setAttribute('maxlength', maxLength);
        updateCounter();
    }
    
    // Add loading state to submit button
    const submitBtn = document.querySelector('button[type="submit"]');
    const submitForm = document.querySelector('form[action="/submit"]');
    if (submitBtn && submitForm) {
        submitForm.addEventListener('submit', function(e) {
            // Check if textarea has content
            const textarea = this.querySelector('textarea[name="regret"]');
            if (!textarea || !textarea.value.trim()) {
                e.preventDefault();
                showNotification('Please write something before submitting!', 'error');
                return;
            }
            
            // Show loading state
            submitBtn.style.opacity = '0.7';
            submitBtn.innerHTML = '📤 Submitting...';
            submitBtn.disabled = true;
        });
    }
    
    // Add fade-in animation to elements
    const observerOptions = {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
    };
    
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.style.opacity = '1';
                entry.target.style.transform = 'translateY(0)';
            }
        });
    }, observerOptions);
    
    // Apply to elements that should animate in
    document.querySelectorAll('.popular-item').forEach(item => {
        item.style.opacity = '0';
        item.style.transform = 'translateY(20px)';
        item.style.transition = 'opacity 0.6s ease, transform 0.6s ease';
        observer.observe(item);
    });
    
    // Copy link functionality
    if (window.location.pathname.startsWith('/r/')) {
        const copyBtn = document.createElement('button');
        copyBtn.textContent = '🔗 Copy Link';
        copyBtn.className = 'btn btn-secondary';
        
        copyBtn.addEventListener('click', async () => {
            try {
                await navigator.clipboard.writeText(window.location.href);
                copyBtn.textContent = '✅ Copied!';
                setTimeout(() => {
                    copyBtn.textContent = '🔗 Copy Link';
                }, 2000);
            } catch (err) {
                console.error('Failed to copy: ', err);
                showNotification('Copy failed. Please copy manually.', 'error');
            }
        });
        
        const navigation = document.querySelector('.navigation');
        if (navigation) {
            navigation.appendChild(copyBtn);
        }
    }
});

// AJAX Like Function
async function likeRegret(regretId) {
    const likeBtn = document.querySelector('.like-button:not(.disabled)');
    const likeCountSpan = document.getElementById('like-count');
    
    if (!likeBtn || likeBtn.classList.contains('disabled')) {
        return;
    }
    
    // Show loading state
    const originalText = likeBtn.innerHTML;
    likeBtn.innerHTML = '⏳ Processing...';
    likeBtn.disabled = true;
    likeBtn.style.opacity = '0.6';
    
    try {
        const response = await fetch(`/like/${regretId}`, {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        });
        
        const data = await response.json();
        
        if (data.success) {
            // Update UI to show liked state
            likeBtn.innerHTML = `❤️ Liked <span class="like-count" id="like-count">${data.likes}</span>`;
            likeBtn.classList.add('disabled');
            likeBtn.style.background = '#95a5a6';
            likeBtn.style.cursor = 'not-allowed';
            likeBtn.style.opacity = '0.7';
            
            // Add pulse animation to the new count
            const newLikeCount = likeBtn.querySelector('.like-count');
            newLikeCount.style.transform = 'scale(1.3)';
            newLikeCount.style.background = 'rgba(255,255,255,0.4)';
            
            setTimeout(() => {
                newLikeCount.style.transform = 'scale(1)';
                newLikeCount.style.background = 'rgba(255,255,255,0.2)';
            }, 300);
            
            // Show success message
            showNotification('❤️ Thank you for sharing the feeling!', 'success');
            
        } else {
            // Show error message
            likeBtn.innerHTML = originalText;
            likeBtn.disabled = false;
            likeBtn.style.opacity = '1';
            showNotification(data.message || 'Unable to like this regret', 'error');
        }
        
    } catch (error) {
        console.error('Error liking regret:', error);
        likeBtn.innerHTML = originalText;
        likeBtn.disabled = false;
        likeBtn.style.opacity = '1';
        showNotification('Network error. Please try again.', 'error');
    }
}

// Notification System
function showNotification(message, type = 'info') {
    // Remove existing notifications
    const existingNotif = document.querySelector('.notification');
    if (existingNotif) {
        existingNotif.remove();
    }
    
    const notification = document.createElement('div');
    notification.className = `notification notification-${type}`;
    notification.textContent = message;
    
    // Styling
    notification.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        padding: 1rem 1.5rem;
        border-radius: 10px;
        color: white;
        font-weight: 500;
        z-index: 1000;
        transform: translateX(100%);
        transition: transform 0.3s ease;
        max-width: 300px;
        word-wrap: break-word;
    `;
    
    // Type-specific styling
    if (type === 'success') {
        notification.style.background = 'linear-gradient(135deg, #4CAF50, #45a049)';
    } else if (type === 'error') {
        notification.style.background = 'linear-gradient(135deg, #f44336, #d32f2f)';
    } else {
        notification.style.background = 'linear-gradient(135deg, #2196F3, #1976D2)';
    }
    
    document.body.appendChild(notification);
    
    // Animate in
    setTimeout(() => {
        notification.style.transform = 'translateX(0)';
    }, 100);
    
    // Auto remove
    setTimeout(() => {
        notification.style.transform = 'translateX(100%)';
        setTimeout(() => {
            notification.remove();
        }, 300);
    }, 3000);
}

// Add some fun easter eggs
let clickCount = 0;
document.addEventListener('click', () => {
    clickCount++;
    if (clickCount === 10) {
        document.body.style.filter = 'hue-rotate(180deg)';
        setTimeout(() => {
            document.body.style.filter = '';
        }, 2000);
        clickCount = 0;
    }
});
