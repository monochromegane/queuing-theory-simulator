import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--path-params',     default='out/params.csv')
parser.add_argument('--path-history',    default='out/history.csv')
parser.add_argument('--path-simulation', default='out/simulation.csv')
parser.add_argument('--path-out',        default='out/plot.png')
args = parser.parse_args()

params      = np.loadtxt(args.path_params,     delimiter=",", skiprows=1)
requests    = np.loadtxt(args.path_history,    delimiter=",", skiprows=1)
simulations = np.loadtxt(args.path_simulation, delimiter=",", skiprows=1)

fig, (ax1, ax2, ax3, ax4) = plt.subplots(4, 1, sharex=True, figsize=(8,12))
_, _, server, lambda_, mu = params
fig.suptitle(r'Simulation of queueing theory(M/M/S) [$S={}$, $\lambda={}$, $\mu={}$, $\rho={:.2f}$]'.format(server, lambda_, mu, lambda_/(mu*server)))

ax1.set_xlim(-5, len(simulations))
for i, req in enumerate(requests):
    arrivalAt, serviceTime, beginAt, endAt = req
    if beginAt == 0.0 and endAt == 0.0:
        continue
    ax1.hlines(y=len(requests)-i, xmin=arrivalAt, xmax=beginAt,    linestyles='dotted', colors='C3', lw=2)
    ax1.hlines(y=len(requests)-i, xmin=beginAt,   xmax=endAt+0.99, linestyles='solid',  colors='C0', lw=2)
ax1.grid(axis='x')
ax1.set_ylabel('Request status')

runnings = simulations[:, 0]
ax2.plot(runnings, color='C0')
ax2.grid()
ax2.set_ylabel('Processing')

waitings = simulations[:, 1]
ax3.plot(waitings, color='C3')
ax3.grid()
ax3.set_ylabel('Waiting')

responseTimes = simulations[:, 2]
ax4.plot(responseTimes, color='C2')
ax4.grid()
ax4.set_ylabel('Response time')

plt.savefig(args.path_out)
