module TargetPlot
import Plots as plt
plt.pyplot()

function f(x,y)
  return ((1.0 - y) * sin(x))^2 + (y * (2.0 - x))^2
end

const plot = function ()
    x = -5.0:0.1:5.0
    y = -5.0:0.1:5.0
    plt.heatmap(x, y, f)
    # plt.plot(x, y, f, st=:wireframe)
end

end