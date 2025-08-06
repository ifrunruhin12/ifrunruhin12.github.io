AOS.init({ duration: 1000, once: true });

// Hero Section Animation Controller
document.addEventListener('DOMContentLoaded', function() {
    // Typewriter effect for greeting
    const typewriter = document.getElementById('typewriter');
    const text = "Hi, I'm";
    let i = 0;
    
    function typeText() {
        if (i < text.length) {
            typewriter.textContent += text.charAt(i);
            i++;
            setTimeout(typeText, 150); // Typing speed
        } else {
            // Hide the cursor after typing is complete
            const cursor = document.querySelector('.cursor');
            if (cursor) {
                cursor.style.display = 'none';
            }
            // Start name animation after typing is complete
            setTimeout(startNameAnimation, 500);
        }
    }
    
    function startNameAnimation() {
        const nameWords = document.querySelectorAll('.name-word');
        nameWords.forEach((word, index) => {
            const delay = parseInt(word.dataset.delay);
            setTimeout(() => {
                word.style.animationDelay = '0s';
                word.style.opacity = '0';
                word.style.animation = 'revealName 0.8s ease-out forwards';
            }, delay);
        });
        
        // Start tagline animation after name animation
        setTimeout(startTaglineAnimation, 1200);
    }
    
    function startTaglineAnimation() {
        const taglineWords = document.querySelectorAll('.tagline-word');
        const taglineSeparators = document.querySelectorAll('.tagline-separator');
        
        // Show the tagline container
        const tagline = document.querySelector('.tagline');
        tagline.style.opacity = '1';
        tagline.style.animation = 'none';
        
        // Animate each word
        taglineWords.forEach((word, index) => {
            const delay = parseInt(word.dataset.delay);
            setTimeout(() => {
                word.style.animationDelay = '0s';
                word.style.animation = 'slideInWord 0.6s ease-out forwards';
            }, delay);
        });
        
        // Animate separators
        taglineSeparators.forEach((separator, index) => {
            setTimeout(() => {
                separator.style.animation = 'fadeIn 0.4s ease-out forwards';
            }, 300 + (index * 200));
        });
    }
    
    // Start the animation sequence
    setTimeout(typeText, 1000); // Wait 1 second before starting
    
    // Add hover effect for name words
    const nameWords = document.querySelectorAll('.name-word');
    nameWords.forEach(word => {
        word.addEventListener('mouseenter', function() {
            this.style.transform = 'translateY(-5px) scale(1.1)';
            this.style.textShadow = '0 0 25px rgba(167, 139, 250, 0.6)';
            this.style.transition = 'all 0.3s ease';
        });
        
        word.addEventListener('mouseleave', function() {
            this.style.transform = 'translateY(0) scale(1)';
            this.style.textShadow = '0 0 20px rgba(167, 139, 250, 0.3)';
        });
    });
    
    // Add subtle floating animation to the hero content
    const heroContent = document.querySelector('.hero-content');
    if (heroContent) {
        heroContent.style.animation = 'float 6s ease-in-out infinite';
    }
    
    // Add floating keyframe animation via CSS
    const style = document.createElement('style');
    style.textContent = `
        @keyframes float {
            0%, 100% { transform: translateY(0px); }
            50% { transform: translateY(-10px); }
        }
    `;
    document.head.appendChild(style);
});

// Add some interactive sparkle effects
function createSparkle(x, y) {
    const sparkle = document.createElement('div');
    sparkle.style.position = 'fixed';
    sparkle.style.left = x + 'px';
    sparkle.style.top = y + 'px';
    sparkle.style.width = '4px';
    sparkle.style.height = '4px';
    sparkle.style.background = '#c084fc';
    sparkle.style.borderRadius = '50%';
    sparkle.style.pointerEvents = 'none';
    sparkle.style.zIndex = '9999';
    sparkle.style.animation = 'sparkle 1s ease-out forwards';
    
    document.body.appendChild(sparkle);
    
    setTimeout(() => {
        sparkle.remove();
    }, 1000);
}

// Add sparkle animation CSS
const sparkleStyle = document.createElement('style');
sparkleStyle.textContent = `
    @keyframes sparkle {
        0% {
            opacity: 1;
            transform: scale(0) rotate(0deg);
        }
        50% {
            opacity: 1;
            transform: scale(1) rotate(180deg);
        }
        100% {
            opacity: 0;
            transform: scale(0) rotate(360deg);
        }
    }
`;
document.head.appendChild(sparkleStyle);

// Add sparkles on hero title click
document.addEventListener('DOMContentLoaded', function() {
    const heroTitle = document.querySelector('.hero-title');
    if (heroTitle) {
        heroTitle.addEventListener('click', function(e) {
            for (let i = 0; i < 8; i++) {
                setTimeout(() => {
                    const rect = this.getBoundingClientRect();
                    const x = rect.left + Math.random() * rect.width;
                    const y = rect.top + Math.random() * rect.height;
                    createSparkle(x, y);
                }, i * 100);
            }
        });
    }
});