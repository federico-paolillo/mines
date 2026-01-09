import { http, HttpResponse } from 'msw'
import { setupWorker } from 'msw/browser'

// This mock worker will intercept requests
const worker = setupWorker(
  http.get('http://localhost:65000/games/123', () => {
    return HttpResponse.json({
        id: '123',
        startTime: Math.floor(Date.now() / 1000), // Started just now
        width: 10,
        height: 10,
        cells: [],
        state: 'Playing',
        lives: 3
    })
  }),
  http.get('http://localhost:65000/games/expired', () => {
    return HttpResponse.json({
        id: 'expired',
        startTime: Math.floor(Date.now() / 1000) - 7200 - 5, // Started 2h 5s ago
        width: 10,
        height: 10,
        cells: [],
        state: 'Playing',
        lives: 3
    })
  })
)

worker.start()
