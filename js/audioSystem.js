// Audio System for Page Flip Sounds
export const audioSystem = {
    audio: null,
    audioPool: [],
    poolSize: 3,
    
    init() {
        try {
            for (let i = 0; i < this.poolSize; i++) {
                const audio = new Audio('assets/page-turn.mp3');
                audio.volume = 0.3;
                audio.preload = 'auto';
                this.audioPool.push(audio);
            }
            this.audio = this.audioPool[0];
            console.log('✅ Page turn audio loaded successfully');
        } catch (error) {
            console.log('Audio not supported, continuing without sound effects');
        }
    },
    
    playPageFlip() {
        if (!this.audioPool.length) return;
        
        let availableAudio = this.audioPool.find(audio => audio.paused || audio.ended);
        if (!availableAudio) {
            availableAudio = this.audioPool[0];
        }
        
        availableAudio.currentTime = 0;
        availableAudio.play().catch(error => {
            console.log('Audio play prevented by browser policy');
        });
    }
};